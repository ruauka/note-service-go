// Package database Package database
package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"web/internal/config"
)

// NewPostgresConnect create connect with DB.
func NewPostgresConnect(cfg *config.Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Postgres.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
