package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"cloud.google.com/go/storage"
	"golang.org/x/net/context"
	"google.golang.org/api/option"

	"github.com/gorilla/mux"
	"github.com/skia-dev/glog"
	"go.skia.org/infra/android_ingest/go/continuous"
	"go.skia.org/infra/android_ingest/go/lookup"
	"go.skia.org/infra/android_ingest/go/parser"
	"go.skia.org/infra/android_ingest/go/recent"
	"go.skia.org/infra/android_ingest/go/upload"
	androidbuildinternal "go.skia.org/infra/go/androidbuildinternal/v2beta1"
	"go.skia.org/infra/go/auth"
	"go.skia.org/infra/go/common"
	"go.skia.org/infra/go/git"
	"go.skia.org/infra/go/httputils"
	"go.skia.org/infra/go/influxdb"
	"go.skia.org/infra/go/login"
	"go.skia.org/infra/go/metrics2"
	"go.skia.org/infra/go/util"
)

// flags
var (
	influxDatabase = flag.String("influxdb_database", influxdb.DEFAULT_DATABASE, "The InfluxDB database.")
	influxHost     = flag.String("influxdb_host", influxdb.DEFAULT_HOST, "The InfluxDB hostname.")
	influxPassword = flag.String("influxdb_password", influxdb.DEFAULT_PASSWORD, "The InfluxDB password.")
	influxUser     = flag.String("influxdb_name", influxdb.DEFAULT_USER, "The InfluxDB username.")
	local          = flag.Bool("local", false, "Running locally if true. As opposed to in production.")
	port           = flag.String("port", ":8000", "HTTP service address (e.g., ':8000')")
	resourcesDir   = flag.String("resources_dir", "", "The directory to find templates, JS, and CSS files. If blank the current directory will be used.")
	workRoot       = flag.String("work_root", "", "Directory location where all the work is done.")
	repoUrl        = flag.String("repo_url", "", "URL of the git repo where buildids are to be stored.")
	branch         = flag.String("branch", "git_master-skia", "The branch where to look for buildids.")
	storageUrl     = flag.String("storage_url", "gs://skia-perf/android-ingest", "The GS URL of where to store the ingested perf data.")
)

var (
	templates      *template.Template
	bucket         *storage.BucketHandle
	gcsPath        string
	converter      *parser.Converter
	process        *continuous.Process
	recentRequests *recent.Recent
	uploads        *metrics2.Counter
	lookupCache    *lookup.Cache
)

func Init() {
	loadTemplates()

	uploads = metrics2.GetCounter("uploads", nil)
	// Create a new auth'd client for androidbuildinternal.
	client, err := auth.NewJWTServiceAccountClient("", "", &http.Transport{Dial: httputils.DialTimeout}, androidbuildinternal.AndroidbuildInternalScope)
	if err != nil {
		glog.Fatalf("Unable to create authenticated client: %s", err)
	}

	if err := os.MkdirAll(*workRoot, 0755); err != nil {
		glog.Fatalf("Failed to create directory %q: %s", *workRoot, err)
	}

	// The repo we're adding commits to.
	checkout, err := git.NewCheckout(*repoUrl, *workRoot)
	if err != nil {
		glog.Fatalf("Unable to create the checkout of %q at %q: %s", *repoUrl, *workRoot, err)
	}
	if err := checkout.Update(); err != nil {
		glog.Fatalf("Unable to update the checkout of %q at %q: %s", *repoUrl, *workRoot, err)
	}

	// checkout isn't go routine safe, but lookup.New() only uses it in New(), so this
	// is safe, i.e. when we later pass checkout to continuous.New().
	lookupCache, err = lookup.New(checkout)
	if err != nil {
		glog.Fatalf("Failed to create buildid lookup cache: %s", err)
	}

	// Start process that adds buildids to the git repo.
	process, err = continuous.New(*branch, checkout, lookupCache, client, *local)
	if err != nil {
		glog.Fatalf("Failed to start continuous process of adding new buildids to git repo: %s", err)
	}
	process.Start()

	var redirectURL = fmt.Sprintf("http://localhost%s/oauth2callback/", *port)
	if !*local {
		redirectURL = "https://android-ingest.skia.org/oauth2callback/"
	}
	if err := login.InitFromMetadataOrJSON(redirectURL, login.DEFAULT_SCOPE, login.DEFAULT_DOMAIN_WHITELIST); err != nil {
		glog.Fatalf("Failed to initialize the login system: %s", err)
	}

	storageHttpClient, err := auth.NewDefaultJWTServiceAccountClient(auth.SCOPE_READ_WRITE)
	if err != nil {
		glog.Fatalf("Problem setting up client OAuth: %s", err)
	}
	storageClient, err := storage.NewClient(context.Background(), option.WithHTTPClient(storageHttpClient))
	if err != nil {
		glog.Fatalf("Problem creating storage client: %s", err)
	}
	gsUrl, err := url.Parse(*storageUrl)
	if err != nil {
		glog.Fatalf("--storage_url value %q is not a valid URL: %s", *storageUrl, err)
	}
	bucket = storageClient.Bucket(gsUrl.Host)
	gcsPath = gsUrl.Path
	if strings.HasPrefix(gcsPath, "/") {
		gcsPath = gcsPath[1:]
	}

	recentRequests = recent.New()

	converter = parser.New(lookupCache, *branch)
}

func makeResourceHandler() func(http.ResponseWriter, *http.Request) {
	fileServer := http.FileServer(http.Dir(*resourcesDir))
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Cache-Control", "max-age=300")
		fileServer.ServeHTTP(w, r)
	}
}

// UploadHandler handles POSTs of images to be analyzed.
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	// Parse incoming JSON.
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		httputils.ReportError(w, r, err, "Failed to read body.")
		return
	}

	// Convert to benchData.
	buf := bytes.NewBuffer(b)
	benchData, err := converter.Convert(buf)
	if err != nil {
		httputils.ReportError(w, r, err, "Failed to find valid incoming JSON.")
		return
	}

	// Write the benchData out as JSON in the right spot in Google Storage.
	writer := bucket.Object(upload.ObjectPath(benchData, gcsPath, time.Now().UTC())).NewWriter(context.Background())
	if err := json.NewEncoder(writer).Encode(benchData); err != nil {
		httputils.ReportError(w, r, err, "Failed to write converted JSON body.")
		return
	}
	util.Close(writer)

	// Store locally.
	recentRequests.Add(b)

	uploads.Inc(1)
}

// IndexContent is the data passed to the index.html template.
type IndexContext struct {
	Recent      []*recent.Request
	LastBuildId int64
}

// MainHandler displays the main page with the last MAX_RECENT Requests.
func MainHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	user := login.LoggedInAs(r)
	if !*local && user == "" {
		http.Redirect(w, r, login.LoginURL(w, r), http.StatusTemporaryRedirect)
		return
	}
	if *local {
		loadTemplates()
	}

	var lastBuildId int64 = -1
	// process is nil when testing.
	if process != nil {
		lastBuildId, _, _, _ = process.Last()
	}

	indexContent := &IndexContext{
		Recent:      recentRequests.List(),
		LastBuildId: lastBuildId,
	}

	if err := templates.ExecuteTemplate(w, "index.html", indexContent); err != nil {
		glog.Errorf("Failed to expand template: %s", err)
	}
}

// redirectHandler handles the links that we added to the git repo and redirects
// them to the source android-build dashboard.
func redirectHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]

	http.Redirect(w, r, fmt.Sprintf("https://android-build.googleplex.com/builds/branches/%s/grid?head=%s&tail=%s", *branch, id, id), http.StatusFound)
}

func loadTemplates() {
	templates = template.Must(template.New("").Delims("{%", "%}").ParseFiles(
		filepath.Join(*resourcesDir, "templates/index.html"),

		// Sub templates used by other templates.
		filepath.Join(*resourcesDir, "templates/header.html"),
	))
}

func main() {
	defer common.LogPanic()
	if *local {
		common.Init()
	} else {
		common.InitWithMetrics2("androidingest", influxHost, influxUser, influxPassword, influxDatabase, local)
	}
	if *workRoot == "" {
		glog.Fatal("The --work_root flag must be supplied.")
	}
	if *repoUrl == "" {
		glog.Fatal("The --repo_url flag must be supplied.")
	}

	Init()

	r := mux.NewRouter()
	r.PathPrefix("/res/").HandlerFunc(makeResourceHandler())
	r.HandleFunc("/upload", UploadHandler)
	r.HandleFunc("/r/{id:[a-zA-Z0-9]+}", redirectHandler)
	r.HandleFunc("/", MainHandler)
	r.HandleFunc("/oauth2callback/", login.OAuth2CallbackHandler)
	r.HandleFunc("/logout/", login.LogoutHandler)
	r.HandleFunc("/loginstatus/", login.StatusHandler)

	http.Handle("/", httputils.LoggingGzipRequestResponse(r))
	glog.Infoln("Ready to serve.")
	glog.Fatal(http.ListenAndServe(*port, nil))
}
