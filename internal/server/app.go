package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/lib/pq"

	"github.com/julienschmidt/httprouter"

	"web/internal/adapters/router/user"
	"web/internal/adapters/storage"
	"web/internal/config"
	"web/internal/domain/services"
	"web/pkg/database/postgres"
)

func Execute() {
	cfg := config.GetConfig()

	db, err := postgres.NewPostgresConnect(cfg)
	if err != nil {
		log.Fatalf("failed to init db: %s", err.Error())
	}

	pgDB := storage.NewStorage(db)
	service := services.NewServices(pgDB)

	router := httprouter.New()
	user.Register(router, service)

	srv := NewServer(cfg.App.Port, router)

	go func() {
		if err := srv.Start(); !errors.Is(err, http.ErrServerClosed) && err != nil {
			log.Printf("error occurred while running http server: %s\n", err.Error())
		}
	}()

	log.Println("Starting server...")

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	log.Println("Shutting down server...")

	if err := srv.Stop(ctx); err != nil {
		log.Printf("error occured on srv shutting down: %s\n", err.Error())
	}

	if err := db.Close(); err != nil {
		log.Printf("error occured on db connection close: %s\n", err.Error())
	}
}
