package codereview

import (
	"context"
	"strconv"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"go.skia.org/infra/go/deepequal"
	"go.skia.org/infra/go/gerrit"
	"go.skia.org/infra/go/gerrit/mocks"
	"go.skia.org/infra/go/httputils"
	"go.skia.org/infra/go/testutils"
	"go.skia.org/infra/go/testutils/unittest"
)

var (
	// http.Client used for testing.
	c = httputils.NewTimeoutClient()
)

func TestSearch(t *testing.T) {
	unittest.SmallTest(t)

	// Mock gerrit.
	g := &mocks.GerritInterface{}
	g.On("Config").Return(gerrit.ConfigChromium).Twice()

	// Mock search call to CQ runs.
	cqChangeID1 := int64(123)
	cqChangeInfo1 := &gerrit.ChangeInfo{Issue: cqChangeID1}
	g.On(
		"Search",
		testutils.AnyContext,
		GerritOpenChangesNum,
		true,
		gerrit.SearchStatus(gerrit.ChangeStatusOpen),
		gerrit.SearchLabel(gerrit.LabelCommitQueue, strconv.Itoa(gerrit.LabelCommitQueueSubmit)),
	).Return([]*gerrit.ChangeInfo{cqChangeInfo1}, nil).Once()
	g.On("GetIssueProperties", testutils.AnyContext, cqChangeID1).Return(cqChangeInfo1, nil).Once()

	// Mock search call to dry-runs.
	dryRunChangeID1 := int64(123) // Same ID as cqChangeID1 to test for deduplication.
	dryRunChangeID2 := int64(345)
	dryRunChangeID3 := int64(902)
	dryRunChangeInfo1 := &gerrit.ChangeInfo{Issue: dryRunChangeID1}
	dryRunChangeInfo2 := &gerrit.ChangeInfo{Issue: dryRunChangeID2}
	dryRunChangeInfo3 := &gerrit.ChangeInfo{Issue: dryRunChangeID3}
	g.On(
		"Search",
		testutils.AnyContext,
		GerritOpenChangesNum,
		true,
		gerrit.SearchStatus(gerrit.ChangeStatusOpen),
		gerrit.SearchLabel(gerrit.LabelCommitQueue, strconv.Itoa(gerrit.LabelCommitQueueDryRun)),
	).Return([]*gerrit.ChangeInfo{dryRunChangeInfo1, dryRunChangeInfo2, dryRunChangeInfo3}, nil)

	cr := gerritCodeReview{
		gerritClient: g,
		cfg:          gerrit.ConfigChromium,
	}
	changes, err := cr.Search(context.Background())
	require.NoError(t, err)
	require.Len(t, changes, 3)
	require.True(t, deepequal.DeepEqual([]*gerrit.ChangeInfo{cqChangeInfo1, dryRunChangeInfo2, dryRunChangeInfo3}, changes))
}

func TestRemoveFromCQ(t *testing.T) {
	unittest.SmallTest(t)

	comment := "SkCQ is no longer looking at this change"
	notifyReason := "SkCQ run failed."
	changeID := int64(123)
	accountID1 := 111111
	accountID2 := 222222
	accountID3 := 333333
	ci := &gerrit.ChangeInfo{
		Issue: changeID,
		Owner: &gerrit.Person{
			Email: "batman@gotham.com",
		},
		Labels: map[string]*gerrit.LabelEntry{
			gerrit.LabelCommitQueue: {
				All: []*gerrit.LabelDetail{
					{
						Value:     gerrit.LabelCommitQueueDryRun,
						AccountID: accountID1,
					},
					{
						Value:     gerrit.LabelCommitQueueNone,
						AccountID: accountID2,
					},
					{
						Value:     gerrit.LabelCommitQueueSubmit,
						AccountID: accountID3,
					},
				},
			},
			// This should be ignored.
			gerrit.LabelCodeReview: {
				All: []*gerrit.LabelDetail{
					{
						Value:     gerrit.LabelCodeReviewApprove,
						AccountID: accountID1,
					},
				},
			},
		},
	}

	// Mock gerrit.
	g := &mocks.GerritInterface{}
	g.On("DeleteVote", testutils.AnyContext, changeID, gerrit.LabelCommitQueue, accountID1, gerrit.NotifyNone, true).Return(nil).Once()
	g.On("DeleteVote", testutils.AnyContext, changeID, gerrit.LabelCommitQueue, accountID3, gerrit.NotifyNone, true).Return(nil).Once()
	g.On("SetReview", testutils.AnyContext, ci, comment, map[string]int{}, []string{}, gerrit.NotifyOwner, mock.Anything, AutogeneratedCommentTag, 0, mock.Anything).Return(nil).Once()

	cr := gerritCodeReview{
		gerritClient: g,
		cfg:          gerrit.ConfigChromium,
	}
	cr.RemoveFromCQ(context.Background(), ci, comment, notifyReason)
}

func TestGetSubmittedTogether(t *testing.T) {
	unittest.SmallTest(t)

	changeID1 := "change1"
	issue1 := int64(123)
	changeID2 := "change2"
	issue2 := int64(324)
	changeID3 := "change3"
	issue3 := int64(523)
	changeInfo1 := &gerrit.ChangeInfo{
		Id:    changeID1,
		Issue: issue1,
	}
	changeInfo2 := &gerrit.ChangeInfo{
		Id:    changeID2,
		Issue: issue2,
	}
	changeInfo3 := &gerrit.ChangeInfo{
		Id:    changeID3,
		Issue: issue3,
	}

	// Mock gerrit with a dependency chain of
	// changeInfo3<-changeInfo2<-changeInfo1. changeInfo1 also has one hidden
	// dependency.
	g := &mocks.GerritInterface{}
	g.On("SubmittedTogether", testutils.AnyContext, changeInfo1).Return([]*gerrit.ChangeInfo{changeInfo1}, 0, nil).Once()
	g.On("SubmittedTogether", testutils.AnyContext, changeInfo2).Return([]*gerrit.ChangeInfo{changeInfo1, changeInfo2}, 0, nil).Once()
	g.On("SubmittedTogether", testutils.AnyContext, changeInfo3).Return([]*gerrit.ChangeInfo{changeInfo1, changeInfo2, changeInfo3}, 1, nil).Once()
	g.On("GetIssueProperties", testutils.AnyContext, issue1).Return(changeInfo1, nil).Once()
	g.On("GetIssueProperties", testutils.AnyContext, issue2).Return(changeInfo2, nil).Once()
	g.On("GetIssueProperties", testutils.AnyContext, issue3).Return(changeInfo3, nil).Once()

	cr := gerritCodeReview{
		gerritClient: g,
		cfg:          gerrit.ConfigChromium,
	}
	// changeInfo1 has no dependencies.
	submittedTogetherChanges1, err1 := cr.GetSubmittedTogether(context.Background(), changeInfo1)
	require.NoError(t, err1)
	require.True(t, deepequal.DeepEqual([]*gerrit.ChangeInfo{}, submittedTogetherChanges1))
	// changeInfo2 has one dependency.
	submittedTogetherChanges2, err2 := cr.GetSubmittedTogether(context.Background(), changeInfo2)
	require.NoError(t, err2)
	require.True(t, deepequal.DeepEqual([]*gerrit.ChangeInfo{changeInfo1}, submittedTogetherChanges2))
	// changeInfo3 has 2 dependencies.
	submittedTogetherChanges3, err3 := cr.GetSubmittedTogether(context.Background(), changeInfo3)
	require.Error(t, err3)
	require.Nil(t, submittedTogetherChanges3)
}

func TestGetEquivalentPatchSetIDs(t *testing.T) {
	unittest.SmallTest(t)

	changeInfo := &gerrit.ChangeInfo{
		// Most recent revisions are first.
		Patchsets: []*gerrit.Revision{
			{
				Number: 1,
				Kind:   gerrit.PatchSetKindCodeChange,
			},
			{
				Number: 2,
				Kind:   gerrit.PatchSetKindTrivialRebase,
			},
			{
				Number: 3,
				Kind:   gerrit.PatchSetKindNoCodeChange,
			},
			{
				Number: 4,
				Kind:   gerrit.PatchSetKindRework,
			},
			{
				Number: 5,
				Kind:   gerrit.PatchSetKindCodeChange,
			},
		},
	}

	cr := gerritCodeReview{}
	// #PS5 has no equivalent patch sets.
	patchsetIDs := cr.GetEquivalentPatchSetIDs(changeInfo, 5)
	require.True(t, deepequal.DeepEqual([]int64{5}, patchsetIDs))
	// #PS4 has no equivalent patch sets.
	patchsetIDs = cr.GetEquivalentPatchSetIDs(changeInfo, 4)
	require.True(t, deepequal.DeepEqual([]int64{4}, patchsetIDs))
	// #PS3 has 2 other equivalent patch sets.
	patchsetIDs = cr.GetEquivalentPatchSetIDs(changeInfo, 3)
	require.True(t, deepequal.DeepEqual([]int64{3, 2, 1}, patchsetIDs))
	// #PS2 has 1 other equivalent patch sets.
	patchsetIDs = cr.GetEquivalentPatchSetIDs(changeInfo, 2)
	require.True(t, deepequal.DeepEqual([]int64{2, 1}, patchsetIDs))
	// #PS1 has no equivalent patch sets.
	patchsetIDs = cr.GetEquivalentPatchSetIDs(changeInfo, 1)
	require.True(t, deepequal.DeepEqual([]int64{1}, patchsetIDs))
	// Test for non-existent patch sets.
	patchsetIDs = cr.GetEquivalentPatchSetIDs(changeInfo, 0)
	require.Len(t, patchsetIDs, 0)
	patchsetIDs = cr.GetEquivalentPatchSetIDs(changeInfo, 6)
	require.Len(t, patchsetIDs, 0)
}

func TestGetChangeRef(t *testing.T) {
	unittest.SmallTest(t)

	changeInfo := &gerrit.ChangeInfo{
		Issue: 443102,
		// Most recent revisions are first.
		Patchsets: []*gerrit.Revision{
			{
				Number: 1,
				Kind:   gerrit.PatchSetKindCodeChange,
			},
			{
				Number: 2,
				Kind:   gerrit.PatchSetKindTrivialRebase,
			},
		},
	}
	cr := gerritCodeReview{}

	require.Equal(t, "refs/changes/02/443102/2", cr.GetChangeRef(changeInfo))
}

func TestGetCQVoters(t *testing.T) {
	unittest.SmallTest(t)

	// Mock gerrit.
	g := &mocks.GerritInterface{}
	g.On("Config").Return(gerrit.ConfigChromium).Twice()

	changeID := int64(123)
	email1 := "batman@gotham.com"
	email2 := "superman@krypton.com"
	email3 := "joker@gotham.com"
	email4 := "pengiun@gotham.com"
	labelDetailWithCQ := []*gerrit.LabelDetail{
		{
			Value: gerrit.LabelCommitQueueDryRun,
			Email: email1,
		},
		{
			Value: gerrit.LabelCommitQueueNone,
			Email: email2,
		},
		{
			Value: gerrit.LabelCommitQueueSubmit,
			Email: email3,
		},
		{
			Value: gerrit.LabelCommitQueueSubmit,
			Email: email4,
		},
	}
	labelDetailWithDryRun := []*gerrit.LabelDetail{
		{
			Value: gerrit.LabelCommitQueueDryRun,
			Email: email1,
		},
		{
			Value: gerrit.LabelCommitQueueNone,
			Email: email2,
		},
	}
	labelDetailWithNone := []*gerrit.LabelDetail{
		{
			Value: gerrit.LabelCommitQueueNone,
			Email: email2,
		},
	}
	labelDetailEmpty := []*gerrit.LabelDetail{}

	changeInfo := &gerrit.ChangeInfo{
		Issue: changeID,
		Labels: map[string]*gerrit.LabelEntry{
			gerrit.LabelCommitQueue: {
				All: labelDetailEmpty,
			},
			// This should be ignored.
			gerrit.LabelCodeReview: {
				All: []*gerrit.LabelDetail{
					{
						Value: gerrit.LabelCodeReviewApprove,
						Email: email1,
					},
				},
			},
		},
	}
	cr := gerritCodeReview{
		gerritClient: g,
		cfg:          gerrit.ConfigChromium,
	}

	tests := []struct {
		labelDetails   []*gerrit.LabelDetail
		expectedVoters []string
	}{
		{
			labelDetails:   labelDetailEmpty,
			expectedVoters: []string{},
		},
		{
			labelDetails:   labelDetailWithNone,
			expectedVoters: []string{},
		},
		{
			labelDetails:   labelDetailWithDryRun,
			expectedVoters: []string{email1},
		},
		{
			labelDetails:   labelDetailWithCQ,
			expectedVoters: []string{email3, email4},
		},
	}

	for _, test := range tests {
		changeInfo.Labels[gerrit.LabelCommitQueue].All = test.labelDetails
		voters := cr.GetCQVoters(context.Background(), changeInfo)
		require.True(t, deepequal.DeepEqual(test.expectedVoters, voters))
	}
}
