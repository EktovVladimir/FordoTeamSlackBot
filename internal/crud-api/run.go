package crud_api

import (
	"context"
	"errors"
	"fmt"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/config"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type CrudApi struct {
	cfg config.CrudApiConfig
}

func New(cfg config.CrudApiConfig) *CrudApi {
	return &CrudApi{cfg}
}

func (a *CrudApi) Start(ctx context.Context) error {
	r := mux.NewRouter()

	r.HandleFunc("/users", getUsers).Methods("GET")

	addr := fmt.Sprintf("%s:%d", a.cfg.Host, a.cfg.Port)

	server := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	logrus.Infof("Crud Api start listen and serve %s", addr)

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logrus.Fatal(err)
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
