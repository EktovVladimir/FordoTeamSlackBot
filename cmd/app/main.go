package main

import (
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/app"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/config"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/environment"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/logger"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"os"
	"os/signal"
	"syscall"
)

const appName = "team_frodo"

// @title						API автоматизации рабочих процессов
// @version					1
// @host						localhost:8080/
//
// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
// @description				Type "Bearer" followed by a space and JWT token.
func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	defer logrus.Info("Application finished")

	environment.InitGlobal()
	cfg := config.Load(appName)
	logger.Init(cfg.Log)

	application := app.New(cfg, appName)

	application.Run(ctx).Wait()
}
