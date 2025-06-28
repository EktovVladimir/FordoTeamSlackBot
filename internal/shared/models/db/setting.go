package db

import (
	"github.com/uptrace/bun"
)

type Setting struct {
	bun.BaseModel   `bun:"table:settings,alias:s"`
	UniqFields      `bun:",embed"`
	AuditableFields `bun:",embed"`

	Key   string `bun:",notnull"`
	Value string `bun:",notnull"`
}
