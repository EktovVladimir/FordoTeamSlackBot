package main

import (
	"context"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/crud-api"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/config"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/db/json_db"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/environment"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/logger"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	defer logrus.Info("Crud API Application finished")

	environment.InitGlobal()
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	logger.Init(cfg.Log)

	store := json_db.NewStore(cfg.JsonStore)
	if err := store.LoadFromFile(); err != nil {
		logrus.Errorf("Error loading json data files: %v", err)
		return
	}

	defer func() {
		if err := store.SyncFiles(); err != nil {
			logrus.Errorf("Failed to sync files: %v", err)
		}
	}()

	store.StartFileSync(ctx, time.Minute*1)

	crudApi := crud_api.New(cfg.CrudApi, store)
	_ = crudApi.Start(ctx)
}
