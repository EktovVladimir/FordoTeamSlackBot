package github

type PullRequest struct {
	Branch     string
	BaseBranch string
	Number     string
	Repo       string
	State      string
}
