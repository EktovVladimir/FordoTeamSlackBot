package db

import "time"

type Deletable interface {
	IsDeleted() bool
}

type DeletedFields struct {
	DeletedAt time.Time `bun:",nullzero"`
}

func (d *DeletedFields) IsDeleted() bool {
	return !d.DeletedAt.IsZero()
}
