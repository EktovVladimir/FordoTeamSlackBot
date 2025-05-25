package git

type PullRequest struct {
	Branch     string
	BaseBranch string
	Number     string
	Repo       string
	State      string
	Requester  User
	Reviewers  []User
}
