package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/handlers/auth_handler"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/handlers/management_handler"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/middlewares"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/environment"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type server struct {
	app               *app
	managementHandler *management_handler.Handler
	authHandler       *auth_handler.Handler
}

func newServer(
	app *app,
	managementHandler *management_handler.Handler,
	authHandler *auth_handler.Handler) *server {
	return &server{app, managementHandler, authHandler}
}

func (s *server) Start(ctx context.Context) {
	gin.SetMode(gin.ReleaseMode)

	if environment.IsDev {
		gin.SetMode(gin.DebugMode)
	}

	router := gin.New()

	router.Use(gin.Recovery())

	authGr := router.Group("/auth")
	s.authHandler.SetupRoutes(authGr)

	router.Use(middlewares.TokenAuthMiddleware(s.app.cfg.Auth))
	router.Use(middlewares.LoggingMiddleware())

	managementGr := router.Group("/management")
	s.managementHandler.SetupRoutes(ctx, managementGr)

	serverConf := s.app.cfg.Server
	addr := fmt.Sprintf("%s:%d", serverConf.Host, serverConf.Port)

	httpServer := &http.Server{
		Addr:         addr,
		Handler:      router,
		ReadTimeout:  serverConf.ReadTimeout,
		WriteTimeout: serverConf.WriteTimeout,
	}

	serverErr := make(chan error, 1)
	go func() {
		logrus.Infof("%s httpServer starting on %s", s.app.appName, addr)
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			serverErr <- fmt.Errorf("failed to start httpServer: %w", err)
		}
	}()

	select {
	case err := <-serverErr:
		logrus.Errorf("%s httpServer has error: %v", s.app.appName, err)
	case <-ctx.Done():
		logrus.Infof("%s httpServer received shutdown signal", s.app.appName)
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(shutdownCtx); err != nil {
		logrus.Errorf("%s server failed to shutdown httpServer gracefully: %v", s.app.appName, err)
		return
	}

	logrus.Infof("%s httpServer stopped gracefully", s.app.appName)
}
