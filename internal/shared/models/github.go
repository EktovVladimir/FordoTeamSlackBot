package models

type PullRequestRef struct {
	Owner  string
	Repo   string
	Number int
}

func NewPullRequestRef(owner string, repo string, number int) *PullRequestRef {
	return &PullRequestRef{owner, repo, number}
}

type CommitInfo struct {
	Message string
}
