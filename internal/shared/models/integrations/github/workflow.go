package github

import (
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/integrations/git"
)

type Workflow struct {
	Event string
	Type  string
}

type PullRequestWorkflow struct {
	Workflow
	PullRequest git.PullRequest
}

type ReviewWorkflow struct {
	Workflow
	PullRequest git.PullRequest
	State       string
}

type RunWorkflow struct {
	Workflow
	Conclusion string
}
