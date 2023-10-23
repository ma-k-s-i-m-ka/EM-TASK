package http

import (
	"EM-TASK/app/internal/core"
	"EM-TASK/app/internal/infrastructure/config"
	"EM-TASK/app/internal/infrastructure/logger"
	"context"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/pkg/browser"
	"net/http"
)

type Server interface {
	Run() error
	Shutdown(ctx context.Context) error
}

type server struct {
	core       core.Core
	router     chi.Router
	httpServer *http.Server
}

func NewHttpServer(core core.Core) (Server, error) {
	router := chi.NewRouter()
	httpPort := config.Config.HTTP.Port
	httpHost := config.Config.HTTP.Host

	s := &server{
		core:   core,
		router: router,
		httpServer: &http.Server{
			Addr:    fmt.Sprintf("%s:%s", httpHost, httpPort),
			Handler: router,
		},
	}

	s.initRouter()

	return s, nil
}

func (s *server) Run() error {
	logger.Logger.Info().Msgf("Starting HTTP server at %s", s.httpServer.Addr)

	err := browser.OpenURL("http://localhost:" + config.Config.HTTP.Port + "/swagger/")
	if err != nil {
		logger.Logger.Info().Msgf("Failed to open documentation in the browser: %s. Try [http://localhost:3005/swagger/index.html]", err)
	}

	err = browser.OpenURL("http://localhost:" + config.Config.HTTP.Port + "/")
	if err != nil {
		logger.Logger.Info().Msgf("Failed to open web-site in the browser: %s. Try [http://localhost:3005/]", err)
	}

	err = s.httpServer.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

func (s *server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
