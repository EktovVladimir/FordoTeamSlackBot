package crud_api

import (
	"context"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"net"
	"net/http"
	"time"
)

func (a *CrudApi) Start(ctx context.Context) {
	routes := a.setupRoutes()
	addr := fmt.Sprintf("%s:%d", a.cfg.Host, a.cfg.Port)

	server := &http.Server{
		Addr:         addr,
		Handler:      routes,
		ReadTimeout:  time.Duration(a.cfg.ReadTimeoutSec) * time.Second,
		WriteTimeout: time.Duration(a.cfg.WriteTimeoutSec) * time.Second,
		BaseContext: func(net.Listener) context.Context {
			return a.contextWithApi(ctx)
		},
	}

	serverErr := make(chan error, 1)
	go func() {
		logrus.Infof("Crud API starting on %s", addr)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			serverErr <- fmt.Errorf("failed to start server: %w", err)
		}
	}()

	select {
	case err := <-serverErr:
		logrus.Errorf("Crud API has error: %v", err)
	case <-ctx.Done():
		logrus.Info("Crud API received shutdown signal")
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		logrus.Errorf("Crud Api failed to shutdown server gracefully: %v", err)
		return
	}

	logrus.Info("Crud Api server stopped gracefully")
}
