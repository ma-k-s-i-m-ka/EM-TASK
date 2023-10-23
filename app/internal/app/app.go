package app

import (
	service "EM-TASK/app/internal"
	"EM-TASK/app/internal/core"
	"EM-TASK/app/internal/infrastructure/config"
	"EM-TASK/app/internal/infrastructure/logger"
	"EM-TASK/app/internal/interfaces/http"
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type application struct {
	core core.Core
}

func New(configPath string) (service.Service, error) {
	err := config.Init(configPath)
	if err != nil {
		return nil, err
	}
	loggerConfig := config.Config.Logger

	logger.Init(&logger.Config{
		LogToConsole:     loggerConfig.LogToConsole,
		EncodeLogsAsJson: loggerConfig.EncodeLogsAsJson,
		LogToFile:        loggerConfig.LogToFile,
		Directory:        loggerConfig.Directory,
		Filename:         loggerConfig.Filename,
		MaxSize:          loggerConfig.MaxSize,
		MaxBackups:       loggerConfig.MaxBackups,
		MaxAge:           loggerConfig.MaxAge,
	})
	logger.Logger.Info().Msg("Logger initialized")
	logger.Logger.Info().Msg("Configuration file loaded")

	c, err := core.New(context.Background())
	if err != nil {
		return nil, err
	}

	return &application{
		core: c,
	}, nil
}

func (app *application) Start() error {
	httpServer, err := http.NewHttpServer(app.core)
	if err != nil {
		return err
	}

	go func() {
		err = httpServer.Run()
		if err != nil {
			logger.Logger.Error().Err(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signals := []os.Signal{syscall.SIGABRT, syscall.SIGQUIT, syscall.SIGHUP, os.Interrupt, syscall.SIGTERM}
	signal.Notify(quit, signals...)

	<-quit
	logger.Logger.Info().Msg("Shutting down the server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = httpServer.Shutdown(ctx); err != nil {
		logger.Logger.Error().Msgf("Server shutdown failed: %v", err)
	}

	logger.Logger.Info().Msg("Server has been shut down")

	return nil
}

func (app *application) Stop(ctx context.Context) error {
	return nil
}
