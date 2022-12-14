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

// NewTestDBClient - create connect with SQLite for mock DB tests. Db in memory. DB Client builder.
func NewTestDBClient() *TestDBClient {
	db, err := sqlx.Open("sqlite3", "file::memory:?cache=shared&_pragma=foreign_keys(1)")
	if err != nil {
		log.Fatalln(err)
	}

	return &TestDBClient{
		Client: db,
	}
}

// SetUp - create a new migrates before tests.
func (t TestDBClient) SetUp() {
	schema := fileOpen("../../../migrations/000001_init.up.sql")
	t.Client.MustExec(strings.ReplaceAll(schema, "serial", "INTEGER"))
}

// TearDown - drop down db after test.
func (t TestDBClient) TearDown() {
	schema := fileOpen("../../../migrations/000001_init.down.sql")
	t.Client.MustExec(schema)
}

// Close - close db connect in unit tests.
func (t TestDBClient) Close() {
	t.Client.Close() //nolint:errcheck,gosec
}

// InsertTestUser - insert test user in DB.
func (t TestDBClient) InsertTestUser(user *model.User) error {
	query := fmt.Sprintf("INSERT INTO %s (username, password) VALUES ($1, $2) RETURNING id", dictionary.UsersTable)
	if err := t.Client.QueryRow(query, user.Username, user.Password).Scan(&user.ID); err != nil {
		return err
	}

	return nil
}

// InsertTestNote - insert test note in DB.
func (t TestDBClient) InsertTestNote(note *model.Note, userID string) error {
	query := fmt.Sprintf("INSERT INTO %s (title, info, user_id) VALUES ($1, $2, $3) RETURNING id", dictionary.NotesTable)
	if err := t.Client.QueryRow(query, note.Title, note.Info, userID).Scan(&note.ID); err != nil {
		return err
	}

	return nil
}

// InsertTestTag - insert test tag in DB.
func (t TestDBClient) InsertTestTag(tag *model.Tag, userID string) error {
	query := fmt.Sprintf("INSERT INTO %s (tagname, user_id) VALUES ($1, $2) RETURNING id", dictionary.TagsTable)
	if err := t.Client.QueryRow(query, tag.TagName, userID).Scan(&tag.ID); err != nil {
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
