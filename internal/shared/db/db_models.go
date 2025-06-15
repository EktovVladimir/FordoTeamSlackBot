package db

type User struct {
	Id         UniqId `json:"id"`
	SlackName  string `json:"slack_name"`
	GithubName string `json:"github_name"`
	Email      string `json:"email"`
}

func (u *User) GetId() UniqId {
	return u.Id
}

func (u *User) SetId(id UniqId) {
	u.Id = id
}

type Deployment struct {
	Id                UniqId `json:"id"`
	ThreadTs          string `json:"thread_ts"`
	PullRequestNumber string `json:"pull_request_number"`
	WorkflowRunId     string `json:"workflow_run_id"`
	Status            string `json:"status"`
}

func (u *Deployment) GetId() UniqId {
	return u.Id
}

func (u *Deployment) SetId(id UniqId) {
	u.Id = id
}

type CodeReview struct {
	Id                UniqId `json:"id"`
	ThreadTs          string `json:"thread_ts"`
	PullRequestNumber string `json:"pull_request_number"`
	Status            string `json:"status"`
}

func (u *CodeReview) GetId() UniqId {
	return u.Id
}

func (u *CodeReview) SetId(id UniqId) {
	u.Id = id
}

type Setting struct {
	Id     UniqId `json:"id"`
	Key    string `json:"key"`
	Source string `json:"source"`
	Value  string `json:"value"`
}

func (u *Setting) GetId() UniqId {
	return u.Id
}

func (u *Setting) SetId(id UniqId) {
	u.Id = id
}
