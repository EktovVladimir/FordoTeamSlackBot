package main

import (
	"context"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/crud-api"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/config"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/environment"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/logger"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	environment.InitGlobal()
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	logger.Init(cfg.Log)

	crudApi := crud_api.New(cfg.CrudApi)
	if err = crudApi.Start(ctx); err != nil {
		logrus.Fatal(err)
	}

	logrus.Info("Crud API Application finished")
}
