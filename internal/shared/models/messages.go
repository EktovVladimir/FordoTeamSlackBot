package models

type CrTreadMessage struct {
	Requester *User
	Reviewers []*User
	Issues    []*JiraIssue
	Prs       []*PullRequestRef
}

type CreatedMessage struct {
	Ts           string
	ChannelId    string
	ResponseText string
	Text         string
}
