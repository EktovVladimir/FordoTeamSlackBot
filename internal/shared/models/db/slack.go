package db

import "github.com/uptrace/bun"

type SlackPostType string

const (
	Thread SlackPostType = "thread"
	Reply  SlackPostType = "reply"
)

type SlackPost struct {
	bun.BaseModel   `bun:"table:slack_post,alias:sp"`
	UniqFields      `bun:",embed"`
	AuditableFields `bun:",embed"`
	DeletedFields   `bun:",embed"`

	ChannelId string         `bun:",notnull"`
	ThreadTs  string         `bun:",notnull"`
	Type      SlackPostType  `bun:",notnull"`
	Meta      map[string]any `bun:",type:jsonb"`
}

func (s SlackPostType) String() string {
	return string(s)
}
