package management_api

import (
	"context"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/config"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/db_adapter/json_db"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/environment"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/logger"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func Run() *sync.WaitGroup {
	wg := &sync.WaitGroup{}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	defer logrus.Info("Management API Application finished")

	environment.InitGlobal()
	cfg := config.Load()
	logger.Init(cfg.Log)

	store := json_db.NewStore(cfg.JsonStore)
	if err := store.LoadFromFile(); err != nil {
		logrus.Errorf("Error loading json data files: %v", err)
		return wg
	}

	defer store.SyncFilesWithLog()

	store.StartFileSync(ctx, 5*time.Minute)

	api := New(cfg.ManagementApi, store)
	api.Start(ctx)

	return wg
}
