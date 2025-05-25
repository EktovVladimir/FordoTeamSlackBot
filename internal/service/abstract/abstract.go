package abstract

import (
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/models/integrations/git"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/models/integrations/issues"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/models/integrations/messenger"
)

type PullRequestService interface {
	Get(repo string, number string) (git.PullRequest, error)
	GetOpened(repo string) ([]git.PullRequest, error)
	GetCommits(repo string, number string) ([]git.Commit, error)
	Create(pr *git.PullRequest) error
}

type IssueTrackerService interface {
	Get(number string) issues.Issue
	GetList(numbers []string) []issues.Issue
}

type MessengerService interface {
	GetThread(channel string, ts string) (messenger.Thread, error)
	SearchThreadLast(channel string, searchCriteria string) (messenger.Thread, error)
	SearchCommentLast(channel string, ts string, searchCriteria string) (messenger.Comment, error)
	CreateThread(channel string, thread *messenger.Thread) error
	UpdateThread(channel string, thread *messenger.Thread) error
	CreateComment(channel string, ts string, comment *messenger.Comment) error
	UpdateComment(channel string, ts string, comment *messenger.Comment) error
}
