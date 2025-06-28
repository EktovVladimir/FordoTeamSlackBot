package db

import (
	"github.com/uptrace/bun"
)

type CodeReview struct {
	bun.BaseModel   `bun:"table:code_reviews,alias:cr"`
	UniqFields      `bun:",embed"`
	AuditableFields `bun:",embed"`

	SlackPostId UniqId     `bun:",notnull"`
	SlackPost   *SlackPost `bun:"rel:belongs-to,join:slack_post_id=id"`

	PullRequests []*PullRequest `bun:"m2m:code_review_pull_requests,join:CodeReview=PullRequest"`

	Status string `bun:",notnull"`

	//Устаревшее
	ThreadTs          string `bun:",notnull"`
	PullRequestNumber string `bun:",notnull"`
}

type CodeReviewToPR struct {
	bun.BaseModel `bun:"table:code_review_pull_requests,alias:crpr"`

	CodeReviewId  UniqId `bun:",notnull,pk"`
	PullRequestId UniqId `bun:",notnull,pk"`

	CodeReview  *CodeReview  `bun:"rel:belongs-to,join:code_review_id=id"`
	PullRequest *PullRequest `bun:"rel:belongs-to,join:pull_request_id=id"`
}
