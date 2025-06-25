package gh_service

import (
	"errors"
	mocks "github.com/EktovVladimir/FordoTeamSlackBot/internal/services/gh_service/mocks"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models"
	"github.com/golang/mock/gomock"
	"github.com/google/go-github/v72/github"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

var (
	errTest = errors.New("test error")
)

func TestGetIssueNumbers_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ts := initTestStruct(ctrl)

	ts.ghClient.EXPECT().
		ListCommits(gomock.Any(), "testOwner", "test-repo", 123, mock.MatchedBy(listOptionsMatcher(1))).
		Return([]*github.RepositoryCommit{
			ghCommitWithMessage("OTAB-123 + add some"),
			ghCommitWithMessage("  otab-777 = message"),
			ghCommitWithMessage("some text OTaB-22 some text"),
			ghCommitWithMessage("OTAC-500 ! fix"),
			ghCommitWithMessage("= no issue commit"),
			ghCommitWithMessage("OTAB-777 = again"),
		}, nil, nil)

	ts.ghClient.EXPECT().
		ListCommits(gomock.Any(), "testOwner", "test-repo", 123, mock.MatchedBy(listOptionsMatcher(2))).
		Return([]*github.RepositoryCommit{}, nil, nil)

	actual, err := ts.service.GetIssueNumbers(
		t.Context(),
		models.NewPullRequestRef("testOwner", "test-repo", 123),
		`OTAB-\d+`)

	expected := []string{
		"OTAB-123",
		"OTAB-777",
		"OTAB-22",
	}

	require.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestGetIssueNumbers_NoMatches_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ts := initTestStruct(ctrl)

	ts.ghClient.EXPECT().
		ListCommits(gomock.Any(), "testOwner", "test-repo", 123, mock.MatchedBy(listOptionsMatcher(1))).
		Return([]*github.RepositoryCommit{
			ghCommitWithMessage("some text OTA-1 some text"),
			ghCommitWithMessage("OTAC-500 ! fix"),
			ghCommitWithMessage("= no issue commit"),
		}, nil, nil)

	ts.ghClient.EXPECT().
		ListCommits(gomock.Any(), "testOwner", "test-repo", 123, mock.MatchedBy(listOptionsMatcher(2))).
		Return([]*github.RepositoryCommit{}, nil, nil)

	actual, err := ts.service.GetIssueNumbers(
		t.Context(),
		models.NewPullRequestRef("testOwner", "test-repo", 123),
		`OTAB-\d+`)

	expected := make([]string, 0)

	require.NoError(t, err)
	assert.Equal(t, expected, actual)
}

type testStruct struct {
	ghClient *mocks.MockgithubPullRequestService
	service  *CommitRetriever
}

func initTestStruct(ctrl *gomock.Controller) *testStruct {
	ghClient := mocks.NewMockgithubPullRequestService(ctrl)

	return &testStruct{
		ghClient: ghClient,
		service:  NewCommitRetriever(ghClient),
	}
}

func listOptionsMatcher(page int) func(*github.ListOptions) bool {
	return func(opt *github.ListOptions) bool {
		return opt.Page == page
	}
}

func ghCommitWithMessage(message string) *github.RepositoryCommit {
	return &github.RepositoryCommit{
		Commit: &github.Commit{
			Message: github.Ptr(message),
		},
	}
}
