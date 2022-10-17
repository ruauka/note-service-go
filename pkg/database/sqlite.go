package database

import (
	"log"
	"os"
	"strings"

	"github.com/jmoiron/sqlx"

	// import sqlite driver.
	_ "github.com/mattn/go-sqlite3"
)

// NewSQLiteConnect - create connect with SQLite for mock DB tests. Db in memory.
func NewSQLiteConnect() *sqlx.DB {
	db, err := sqlx.Open("sqlite3", ":memory:?cache=shared&_pragma=foreign_keys(1)")
	if err != nil {
		log.Fatalln(err)
	}

	return db
}

// SetUp - create a new migrates before tests.
func SetUp(db *sqlx.DB, schema string) {
	db.MustExec(strings.ReplaceAll(schema, "serial", "INTEGER"))
}

// TearDown - drop down db after test.
func TearDown(db *sqlx.DB, schema string) {
	db.MustExec(schema)
}

// FileOpen - open file .sql with migrations.
func FileOpen(path string) string {
	schema, err := os.ReadFile(path) //nolint:gosec
	if err != nil {
		log.Fatalln(err)
	}

	return string(schema)
}
