package crud_api

import (
	"context"
	"errors"
	"fmt"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/api/middleware"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/config"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/db"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type store interface {
	GetUsers() *db.Entity[*db.User]
	GetSettings() *db.Entity[*db.Setting]
	GetDeployments() *db.Entity[*db.Deployment]
	GetCodeReviews() *db.Entity[*db.CodeReview]
}

type CrudApi struct {
	cfg config.CrudApiConfig
	db  store
}

func New(cfg config.CrudApiConfig, db store) *CrudApi {
	return &CrudApi{cfg, db}
}

func (a *CrudApi) Start(ctx context.Context) error {
	r := mux.NewRouter()

	r.Use(middleware.LoggingMiddleware)
	r.HandleFunc("/users", a.getUsers).Methods("GET")
	r.HandleFunc("/users", a.createUser).Methods("POST")

	addr := fmt.Sprintf("%s:%d", a.cfg.Host, a.cfg.Port)

	server := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	logrus.Infof("Crud Api start listen and serve %s", addr)

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logrus.Errorf("Crud Api listen and serve error %v", err)
		}
	}()

	<-ctx.Done()

	if err := server.Shutdown(ctx); err != nil {
		logrus.Errorf("Crud Api failed to shutdown server: %v", err)
		return err
	}

	logrus.Info("Crud Api server shutdown")

	return nil
}
