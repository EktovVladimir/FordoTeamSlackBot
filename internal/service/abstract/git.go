package abstract

import "github.com/EktovVladimir/FordoTeamSlackBot/internal/models/integrations/git"

type PullRequestService interface {
	Get(repo string, number string) (git.PullRequest, error)
	GetOpened(repo string) ([]git.PullRequest, error)
	GetCommits(repo string, number string) ([]git.Commit, error)
	Create(pr *git.PullRequest) error
}
