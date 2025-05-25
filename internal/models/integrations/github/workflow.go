package github

type Workflow struct {
	Event string
	Type  string
}

type PullRequestWorkflow struct {
	Workflow
	PullRequest PullRequest
}

type ReviewWorkflow struct {
	Workflow
	PullRequest PullRequest
	State       string
}

type RunWorkflow struct {
	Workflow
	Conclusion string
}
