package abstract

import (
	git2 "github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/integrations/git"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/integrations/issues"
	messenger2 "github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/integrations/messenger"
)

type PullRequestService interface {
	Get(repo string, number string) (git2.PullRequest, error)
	GetOpened(repo string) ([]git2.PullRequest, error)
	GetCommits(repo string, number string) ([]git2.Commit, error)
	Create(pr *git2.PullRequest) error
}

type IssueTrackerService interface {
	Get(number string) (issues.Issue, error)
	GetList(numbers []string) ([]issues.Issue, error)
}

type MessengerService interface {
	GetUser(email string) (messenger2.User, error)
	GetThread(channel string, ts string) (messenger2.Thread, error)
	SearchThreadLast(channel string, searchCriteria string) (messenger2.Thread, error)
	SearchCommentLast(channel string, ts string, searchCriteria string) (messenger2.Comment, error)
	CreateThread(channel string, thread *messenger2.Thread) error
	UpdateThread(channel string, thread *messenger2.Thread) error
	CreateComment(channel string, ts string, comment *messenger2.Comment) error
	UpdateComment(channel string, ts string, comment *messenger2.Comment) error
}

type SettingsService interface {
	GetStringItem(source Source, key string) string
}

type DeployService interface {
	CreateThread(pr git2.PullRequest, source Source) error
}
