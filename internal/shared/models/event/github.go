package event

import "fmt"

type GithubPullRequestEvent struct {
	Number     string
	Repo       string
	BaseBranch string
	Branch     string
	Action     string
	Requester  string
}

func (e GithubPullRequestEvent) GetInitiator() string {
	return e.Requester
}

func (e GithubPullRequestEvent) Print() string {
	return fmt.Sprintf("pull request event (%s): %s/%s %s->%s from %s",
		e.Action, e.Repo, e.Number, e.Branch, e.BaseBranch, e.Requester)
}
