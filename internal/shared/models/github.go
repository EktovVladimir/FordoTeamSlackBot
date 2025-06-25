package models

import "fmt"

type PullRequestRef struct {
	Owner  string
	Repo   string
	Number int
}

func NewPullRequestRef(owner string, repo string, number int) *PullRequestRef {
	return &PullRequestRef{owner, repo, number}
}

func (p *PullRequestRef) GetUrl() string {
	return fmt.Sprintf("https://github.com/%s/%s/pull/%d", p.Owner, p.Repo, p.Number)
}

type CommitInfo struct {
	Message string
}

type GhReviewer struct {
	Email string
	Login string
}
