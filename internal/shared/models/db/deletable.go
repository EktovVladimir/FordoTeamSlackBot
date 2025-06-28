package db

import "time"

type Deletable interface {
	IsDeleted() bool
}

type DeletedFields struct {
	DeletedAt time.Time `json:"deleted_at" bson:"deleted_at"`
}

func (d *DeletedFields) IsDeleted() bool {
	return !d.DeletedAt.IsZero()
}
