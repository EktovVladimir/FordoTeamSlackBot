package main

import (
	"context"
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
	cfg, err := config.Load(ctx)
	if err != nil {
		log.Fatal(err)
		return
	}

	logrus.Warn(cfg)

	logger.Init(cfg.Log)

	logrus.Info("Application finished")
}
