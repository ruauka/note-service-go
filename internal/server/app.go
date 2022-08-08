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

	_ "github.com/lib/pq"

	"github.com/julienschmidt/httprouter"

	"web/docs"
	"web/internal/adapters/router/handlers/note"
	"web/internal/adapters/router/handlers/tag"
	"web/internal/adapters/router/handlers/user"
	"web/internal/adapters/router/swagger"
	s "web/internal/adapters/storage"
	"web/internal/config"
	"web/internal/domain/services"
	"web/pkg/database/postgres"
	l "web/pkg/logger"
)

func Execute() {
	docs.BuildInfo.Print()

	cfg := config.GetConfig()

	logger := l.NewLogger(cfg)
	loggingMiddleware := l.NewLoggerMiddleware(logger)

	db, err := postgres.NewPostgresConnect(cfg)
	if err != nil {
		logger.Fatal(fmt.Sprintf("failed to init db: %s", err.Error()))
	}

	router := httprouter.New()
	storage := s.NewStorages(db)
	service := services.NewServices(storage)

	swagger.Register(router)
	user.Register(router, service, loggingMiddleware)
	note.Register(router, service, loggingMiddleware)
	tag.Register(router, service, loggingMiddleware)

	srv := NewServer(cfg, router)

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

	if err := srv.Stop(ctx); err != nil {
		logger.Error(fmt.Sprintf("error occured on srv shutting down: %s\n", err.Error()))
	}

	if err := db.Close(); err != nil {
		logger.Error(fmt.Sprintf("error occured on db connection close: %s\n", err.Error()))
	}
}
