package abstract

import (
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/db"
)

type DbRepository interface {
	ConfigRepository
	MessageRepository
}

type ConfigRepository interface {
	GetItem(source string, key string) (db.Setting, error)
	SetItem(item db.Setting) error
}

type MessageRepository interface {
}
