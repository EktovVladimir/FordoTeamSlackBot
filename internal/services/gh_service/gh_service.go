//go:generate mockgen -destination=mocks/mock.go -source gh_service.go

package gh_service

import (
	"github.com/google/go-github/v72/github"
	"golang.org/x/net/context"
)

const (
	pageSize = 100
)

type githubPullRequestService interface {
	ListCommits(context.Context, string, string, int, *github.ListOptions) ([]*github.RepositoryCommit, *github.Response, error)
	ListReviewers(context.Context, string, string, int, *github.ListOptions) (*github.Reviewers, *github.Response, error)
}

type service struct {
	//ghClient *github.Client
	ghClient githubPullRequestService
}
