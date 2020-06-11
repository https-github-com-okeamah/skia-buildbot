// dryrun allows testing an Alert and seeing the regression it would find.
package dryrun

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"go.skia.org/infra/go/auditlog"
	"go.skia.org/infra/go/httputils"
	"go.skia.org/infra/go/metrics2"
	"go.skia.org/infra/go/sklog"
	"go.skia.org/infra/perf/go/alerts"
	"go.skia.org/infra/perf/go/cid"
	"go.skia.org/infra/perf/go/dataframe"
	perfgit "go.skia.org/infra/perf/go/git"
	"go.skia.org/infra/perf/go/regression"
	"go.skia.org/infra/perf/go/shortcut"
	"go.skia.org/infra/perf/go/types"
)

const (
	CLEANUP_DURATION = 5 * time.Minute
)

// UIDomain is the time domain over which to look for commit.
//
// Generated by the <domain-picker-sk> element.
type UIDomain struct {
	Begin       int                   `json:"begin"`       // Beginning of time range in Unix timestamp seconds.
	End         int                   `json:"end"`         // End of time range in Unix timestamp seconds.
	NumCommits  int32                 `json:"num_commits"` // If RequestType is REQUEST_COMPACT, then the number of commits to show before End, and Begin is ignored.
	RequestType dataframe.RequestType `json:"request_type"`
}

// StartRequest is the data POSTed to StartHandler.
type StartRequest struct {
	Config alerts.Alert `json:"config"`
	Domain UIDomain     `json:"domain"`
}

// Id is used to identify StartRequests as stored in Requests.inFlight.
func (s *StartRequest) Id() string {
	return fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%#v", *s))))
}

// StartResponse is the JSON response sent from StartHandler.
type StartResponse struct {
	ID string `json:"id"`
}

// Running is the data stored for each running dryrun.
type Running struct {
	mutex        sync.Mutex
	whenFinished time.Time
	Finished     bool                              `json:"finished"`    // True if the dry run is complete.
	Message      string                            `json:"message"`     // Human readable string describing the dry run state.
	Regressions  map[string]*regression.Regression `json:"regressions"` // All the regressions found so far.
}

// Requests handles HTTP request for doing dryruns.
type Requests struct {
	cidl           *cid.CommitIDLookup
	dfBuilder      dataframe.DataFrameBuilder
	perfGit        *perfgit.Git
	paramsProvider regression.ParamsetProvider // TODO build the paramset from dfBuilder.
	shortcutStore  shortcut.Store
	mutex          sync.Mutex
	inFlight       map[string]*Running
}

// New create a new dryrun Request processor.
func New(cidl *cid.CommitIDLookup, dfBuilder dataframe.DataFrameBuilder, shortcutStore shortcut.Store, paramsProvider regression.ParamsetProvider, perfGit *perfgit.Git) *Requests {
	ret := &Requests{
		cidl:           cidl,
		dfBuilder:      dfBuilder,
		paramsProvider: paramsProvider,
		shortcutStore:  shortcutStore,
		perfGit:        perfGit,
		inFlight:       map[string]*Running{},
	}
	// Start a go routine to clean up old dry runs.
	go ret.cleaner()
	return ret
}

// cleanerStep does a single step of cleaner().
func (d *Requests) cleanerStep() {
	cutoff := time.Now().Add(-CLEANUP_DURATION)
	d.mutex.Lock()
	defer d.mutex.Unlock()
	for k, v := range d.inFlight {
		if !v.Finished {
			continue
		}
		if v.whenFinished.Before(cutoff) {
			delete(d.inFlight, k)
		}
	}
	metrics2.GetInt64Metric("dryrun_inflight", nil).Update(int64(len(d.inFlight)))
}

// cleaner removes old dry runs from inFlight.
func (d *Requests) cleaner() {
	for range time.Tick(CLEANUP_DURATION) {
		d.cleanerStep()
	}
}

func (d *Requests) StartHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req StartRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httputils.ReportError(w, err, "Could not decode POST body.", http.StatusInternalServerError)
		return
	}
	auditlog.Log(r, "dryrun", req)
	if req.Config.Query == "" {
		httputils.ReportError(w, fmt.Errorf("Query was empty."), "A Query is required.", http.StatusInternalServerError)
		return
	}
	if err := req.Config.Validate(); err != nil {
		httputils.ReportError(w, err, "Invalid Alert config.", http.StatusInternalServerError)
		return
	}
	d.mutex.Lock()
	defer d.mutex.Unlock()
	id := req.Id()
	if p, ok := d.inFlight[id]; ok {
		p.mutex.Lock()
		defer p.mutex.Unlock()
		if p.Finished {
			delete(d.inFlight, id)
		}
	}
	if _, ok := d.inFlight[id]; !ok {
		running := &Running{
			Finished:    false,
			Message:     "Starting dry run.",
			Regressions: map[string]*regression.Regression{},
		}
		d.inFlight[id] = running
		go func() {
			ctx := context.Background()
			// Create a callback that will be passed each found Regression.
			cb := func(queryRequest *regression.RegressionDetectionRequest, clusterResponse []*regression.RegressionDetectionResponse, message string) {
				running.mutex.Lock()
				defer running.mutex.Unlock()
				// Loop over clusterResponse, convert each one to a regression, and merge with running.Regressions.
				for _, cr := range clusterResponse {
					c, reg, err := regression.RegressionFromClusterResponse(ctx, cr, &req.Config, d.cidl)
					if err != nil {
						running.Message = "Failed to convert to Regression, some data may be missing."
						sklog.Errorf("Failed to convert to Regression: %s", err)
						return
					}
					id := c.ID()
					running.Message = fmt.Sprintf("Step: %d/%d\nQuery: %q\nLooking for regressions in query results.\n  Commit: %d\n  Details: %q", queryRequest.Step+1, queryRequest.TotalQueries, queryRequest.Query, c.CommitID.Offset, message)
					// We might not have found any regressions.
					if reg.Low == nil && reg.High == nil {
						continue
					}
					if origReg, ok := running.Regressions[id]; !ok {
						running.Regressions[id] = reg
					} else {
						running.Regressions[id] = origReg.Merge(reg)
					}
				}
			}
			progressCallback := func(message string) {
				running.mutex.Lock()
				defer running.mutex.Unlock()
				running.Message = message
			}
			domain := domainFromUIDomain(req.Domain)
			regression.RegressionsForAlert(ctx, &req.Config, domain, d.paramsProvider(), d.shortcutStore, cb, d.perfGit, d.cidl, d.dfBuilder, progressCallback)
			running.mutex.Lock()
			defer running.mutex.Unlock()
			running.Finished = true
			running.whenFinished = time.Now()
			running.Message = running.Message + "\nDry run complete."
		}()
	}
	resp := StartResponse{
		ID: id,
	}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		sklog.Errorf("Failed to encode paramset: %s", err)
	}
}

// domainFromUIDomain converts the UIDomain that domain-picker-sk returns into
// a types.Domain.
func domainFromUIDomain(uiDomain UIDomain) types.Domain {
	return types.Domain{
		N:   uiDomain.NumCommits,
		End: time.Unix(int64(uiDomain.End), 0),
	}
}

// RegressionRow is a Regression found for a specific commit.
type RegressionRow struct {
	CID        *cid.CommitDetail      `json:"cid"`
	Regression *regression.Regression `json:"regression"`
}

// Status is the JSON response sent from StatusHandler.
type Status struct {
	Finished    bool             `json:"finished"`
	Message     string           `json:"message"`
	Regressions []*RegressionRow `json:"regressions"`
}

func (d *Requests) StatusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]

	d.mutex.Lock()
	defer d.mutex.Unlock()

	// Grab the running dryrun.
	running, ok := d.inFlight[id]
	if !ok {
		httputils.ReportError(w, fmt.Errorf("Invalid id: %q", id), "Invalid or expired dry run.", http.StatusInternalServerError)
		return
	}

	status := &Status{
		Finished:    running.Finished,
		Message:     running.Message,
		Regressions: []*RegressionRow{},
	}

	// Convert the Running.Regressions into a properly formed Status response.
	if running.Finished {
		running.mutex.Lock()
		defer running.mutex.Unlock()
		keys := []string{}
		for id := range running.Regressions {
			keys = append(keys, id)
		}
		sort.Strings(keys)

		cids := []*cid.CommitID{}
		for _, key := range keys {
			commitId, err := cid.FromID(key)
			if err != nil {
				httputils.ReportError(w, err, "Failed to parse commit id.", http.StatusInternalServerError)
				return
			}
			cids = append(cids, commitId)
		}

		cidd, err := d.cidl.Lookup(r.Context(), cids)
		if err != nil {
			httputils.ReportError(w, err, "Failed to find commit ids.", http.StatusInternalServerError)
			return
		}
		for _, details := range cidd {
			status.Regressions = append(status.Regressions, &RegressionRow{
				CID:        details,
				Regression: running.Regressions[details.ID()],
			})
		}
	}
	if err := json.NewEncoder(w).Encode(status); err != nil {
		sklog.Errorf("Failed to encode paramset: %s", err)
	}
}
