package git

type PullRequest struct {
	Branch     string
	BaseBranch string
	Number     string
	Repo       string
	Owner      string
	State      string
	Requester  User
	Reviewers  []User
	Title      string
	Labels     []string
}
