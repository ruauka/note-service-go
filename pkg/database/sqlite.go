package database

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jmoiron/sqlx"

	// import sqlite driver.
	_ "github.com/mattn/go-sqlite3"

	"web/internal/domain/entities/model"
	"web/internal/utils/dictionary"
)

// TestDBClient - db Client for unit testing.
type TestDBClient struct {
	Client *sqlx.DB
}

// NewTestDBClient - db Client builder.
func NewTestDBClient() *TestDBClient {
	return &TestDBClient{
		Client: NewSQLiteConnect(),
	}
}

// NewSQLiteConnect - create connect with SQLite for mock DB tests. Db in memory.
func NewSQLiteConnect() *sqlx.DB {
	db, err := sqlx.Open("sqlite3", ":memory:?cache=shared&_pragma=foreign_keys(1)")
	if err != nil {
		log.Fatalln(err)
	}

	return db
}

// SetUp - create a new migrates before tests.
func (t TestDBClient) SetUp() {
	schema := fileOpen("../../../migrate/000001_init.up.sql")
	t.Client.MustExec(strings.ReplaceAll(schema, "serial", "INTEGER"))
}

// TearDown - drop down db after test.
func (t TestDBClient) TearDown() {
	schema := fileOpen("../../../migrate/000001_init.down.sql")
	t.Client.MustExec(schema)
}

// Close - close db connect in unit tests.
func (t TestDBClient) Close() {
	t.Client.Close() //nolint:errcheck,gosec
}

// InsertTestUser - insert user in DB.
func (t TestDBClient) InsertTestUser(user *model.User) error {
	query := fmt.Sprintf("INSERT INTO %s (username, password) VALUES ($1, $2) RETURNING id", dictionary.UsersTable)
	if err := t.Client.QueryRow(query, user.Username, user.Password).Scan(&user.ID); err != nil {
		return err
	}

	return nil
}

// fileOpen - open file .sql with migrations.
func fileOpen(path string) string {
	schema, err := os.ReadFile(path) //nolint:gosec
	if err != nil {
		log.Fatalln(err)
	}

	return string(schema)
}
