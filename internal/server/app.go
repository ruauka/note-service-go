// Package server Package server
package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	// import postgres driver.
	_ "github.com/lib/pq"

	"github.com/julienschmidt/httprouter"

	"web/api"
	"web/internal/adapters/router/handlers/note"
	"web/internal/adapters/router/handlers/tag"
	"web/internal/adapters/router/handlers/user"
	"web/internal/adapters/router/swagger"
	s "web/internal/adapters/storage"
	"web/internal/config"
	"web/internal/domain/services"
	"web/pkg/database"
	l "web/pkg/logger"
)

// Execute main service func.
func Execute() {
	// print service description
	api.BuildInfo.Print()
	// config create
	cfg := config.GetConfig()
	// logger create
	logger := l.NewLogger(cfg)
	loggingMiddleware := l.NewLoggerMiddleware(logger)
	// conn to DB create
	db, err := database.NewPostgresConnect(cfg)
	if err != nil {
		logger.Fatal(fmt.Sprintf("failed to init db: %s", err.Error()))
	}
	// router create
	router := httprouter.New()
	// storage create
	storage := s.NewStorages(db)
	// services (usecases) create
	service := services.NewServices(storage)
	// swagger handler register
	swagger.Register(router)
	// service handlers register
	user.Register(router, service, loggingMiddleware)
	note.Register(router, service, loggingMiddleware)
	tag.Register(router, service, loggingMiddleware)
	// server create
	srv := NewServer(cfg, router)
	// server start
	go func() {
		if err := srv.Start(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Fatal(fmt.Sprintf("error occurred while running http server: %s\n", err.Error()))
		}
	}()

	logger.Info("Starting server...")

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	logger.Info("Shutting down server...")
	// server stop
	if err := srv.Stop(ctx); err != nil {
		logger.Error(fmt.Sprintf("error occurred on srv shutting down: %s\n", err.Error()))
	}
	// close connection with DB
	if err := db.Close(); err != nil {
		logger.Error(fmt.Sprintf("error occurred on db connection close: %s\n", err.Error()))
	}
}
