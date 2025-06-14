package run

import (
	"context"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/db"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/service/store_logger"
	"log"
	"sync"
	"time"
)

func RunHw20(ctx context.Context) *sync.WaitGroup {
	wg := &sync.WaitGroup{}

	db := db.NewFileMemoryDb("file_db")

	if err := db.LoadAll(); err != nil {
		log.Fatal(err)
		return wg
	}

	storeLogger := store_logger.NewStoreLogger(db, 200*time.Millisecond)
	storeLogger.StartGo(ctx)

	return wg
}
