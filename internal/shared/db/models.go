package db

import "time"

type UniqId = int64

type UniqEntity struct {
	Id UniqId `json:"id"`
}

type CreatedEntity struct {
	CreatedDate time.Time `json:"created_date"`
	CreatedBy   UniqId    `json:"created_by"`
}

type UpdatedEntity struct {
	UpdatedDate time.Time `json:"updated_date"`
	UpdatedBy   UniqId    `json:"updated_by"`
}

type User struct {
	UniqEntity
	SlackName  string `json:"slack_name"`
	GithubName string `json:"github_name"`
	Email      string `json:"email"`
}

type Deployment struct {
	UniqEntity
	CreatedEntity
	ThreadTs          string `json:"thread_ts"`
	PullRequestNumber string `json:"pull_request_number"`
	WorkflowRunId     string `json:"workflow_run_id"`
	Status            string `json:"status"`
}

type CodeReview struct {
	UniqEntity
	CreatedEntity
	ThreadTs          string `json:"thread_ts"`
	PullRequestNumber string `json:"pull_request_number"`
	Status            string `json:"status"`
}

type Setting struct {
	UniqEntity
	CreatedEntity
	UpdatedEntity
	Key    string `json:"key"`
	Source string `json:"source"`
	Value  string `json:"value"`
}
