package db

import "time"

type Auditable interface {
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
}

type AuditableFields struct {
	CreatedAt time.Time `json:"created_at" bson:"created_at" bun:",notnull,default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at" bun:",notnull,default:current_timestamp"`
}

func (a AuditableFields) GetCreatedAt() time.Time {
	return a.CreatedAt
}

func (a AuditableFields) GetUpdatedAt() time.Time {
	return a.UpdatedAt
}
