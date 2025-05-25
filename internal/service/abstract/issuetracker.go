package abstract

import "github.com/EktovVladimir/FordoTeamSlackBot/internal/models/integrations/issues"

type IssueTrackerService interface {
	Get(number string) issues.Issue
	GetList(numbers []string) []issues.Issue
}
