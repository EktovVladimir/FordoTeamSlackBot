package db

import "github.com/uptrace/bun"

type SlackPostType string

const (
	Thread SlackPostType = "thread"
	Reply  SlackPostType = "reply"
)

type SlackPost struct {
	bun.BaseModel   `bun:"table:slack_post,alias:sp" bson:"-"`
	UniqFields      `bson:",inline" bun:",embed"`
	AuditableFields `bson:",inline" bun:",embed"`
	DeletedFields   `bson:",inline" bun:",embed"`

	ChannelId string        `json:"channel_id" bson:"channel_id" bun:",notnull"`
	ThreadTs  string        `json:"thread_ts" bson:"thread_ts" bun:",notnull"`
	Type      SlackPostType `json:"type" bson:"type" bun:",notnull"`
}

func (s SlackPostType) String() string {
	return string(s)
}
