package storage

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"

	"web/internal/domain/entities/dto"
	"web/internal/domain/entities/model"
	"web/pkg/database"
)

func TestNoteStorage_CreateNote(t *testing.T) {
	db := database.NewTestDBClient()
	defer db.Close()

	testTable := []struct {
		note, expected *model.Note
		err            error
		testName       string
	}{
		{
			note: &model.Note{
				Title: "test_title",
				Info:  "test_info",
			},
			expected: &model.Note{
				ID:    "1",
				Title: "test_title",
				Info:  "test_info",
			},
			err:      nil,
			testName: "Test-1-OK",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.testName, func(t *testing.T) {
			db.SetUp()
			defer db.TearDown()

			storage := NewNoteStorage(db.Client)

			actual, err := storage.CreateNote(testCase.note, "1")
			if err != nil {
				log.Fatalln(err.Error())
			}

			require.Equal(t, testCase.expected, actual)
			require.NoError(t, testCase.err, err)
		})
	}
}

func TestNoteStorage_GetNoteByID(t *testing.T) {
	db := database.NewTestDBClient()
	defer db.Close()

	testTable := []struct {
		note     *model.Note
		expected *dto.NoteResp
		err      error
		testName string
	}{
		{
			note: &model.Note{
				Title: "test_title",
				Info:  "test_info",
			},
			expected: &dto.NoteResp{
				Title: "test_title",
				Info:  "test_info",
			},
			testName: "Test-1-OK",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.testName, func(t *testing.T) {
			db.SetUp()
			defer db.TearDown()

			if err := db.InsertTestNote(testCase.note, "1"); err != nil {
				log.Fatalln(err.Error())
			}

			storage := NewNoteStorage(db.Client)

			actual, err := storage.GetNoteByID("1", "1")
			if err != nil {
				log.Fatalln(err.Error())
			}

			require.Equal(t, testCase.expected, actual)
			require.NoError(t, testCase.err, err)
		})
	}
}

func TestNoteStorage_GetAllNotesByUser(t *testing.T) {
	db := database.NewTestDBClient()
	defer db.Close()

	testTable := []struct {
		note     *model.Note
		expected []dto.NotesResp
		err      error
		testName string
	}{
		{
			note: &model.Note{
				Title: "test_title",
				Info:  "test_info",
			},
			expected: []dto.NotesResp{
				{
					ID:    "1",
					Title: "test_title",
					Info:  "test_info",
				},
			},
			testName: "Test-1-OK",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.testName, func(t *testing.T) {
			db.SetUp()
			defer db.TearDown()

			if err := db.InsertTestNote(testCase.note, "1"); err != nil {
				log.Fatalln(err.Error())
			}

			storage := NewNoteStorage(db.Client)

			actual, err := storage.GetAllNotesByUser("1")
			if err != nil {
				log.Fatalln(err.Error())
			}

			require.Equal(t, testCase.expected, actual)
			require.NoError(t, testCase.err, err)
		})
	}
}

func TestNoteStorage_UpdateNote(t *testing.T) {
	db := database.NewTestDBClient()
	defer db.Close()

	noteTitle := "note_title1"
	noteInfo := "note_info1"

	testTable := []struct {
		note     *model.Note
		newNote  *dto.NoteUpdate
		expected error
		testName string
	}{
		{
			note: &model.Note{
				Title: "test_title",
				Info:  "test_info",
			},
			newNote: &dto.NoteUpdate{
				Title: &noteTitle,
				Info:  &noteInfo,
			},
			expected: nil,
			testName: "Test-1-OK",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.testName, func(t *testing.T) {
			db.SetUp()
			defer db.TearDown()

			if err := db.InsertTestNote(testCase.note, "1"); err != nil {
				log.Fatalln(err.Error())
			}

			storage := NewNoteStorage(db.Client)

			actual := storage.UpdateNote(testCase.newNote, "1")

			require.NoError(t, testCase.expected, actual)
		})
	}
}

func TestNoteStorage_DeleteNote(t *testing.T) {
	db := database.NewTestDBClient()
	defer db.Close()

	testTable := []struct {
		note     *model.Note
		expected int
		err      error
		testName string
	}{
		{
			note: &model.Note{
				Title: "test_title",
				Info:  "test_info",
			},
			expected: 1,
			err:      nil,
			testName: "Test-1-OK",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.testName, func(t *testing.T) {
			db.SetUp()
			defer db.TearDown()

			if err := db.InsertTestNote(testCase.note, "1"); err != nil {
				log.Fatalln(err.Error())
			}

			storage := NewNoteStorage(db.Client)

			actual, err := storage.DeleteNote("1", "1")
			if err != nil {
				log.Fatalln(err.Error())
			}

			require.Equal(t, testCase.expected, actual)
			require.NoError(t, testCase.err, err)
		})
	}
}

func TestNoteStorage_SetTags(t *testing.T) {
	db := database.NewTestDBClient()
	defer db.Close()

	testTable := []struct {
		note     *model.Note
		tags     map[string]string
		expected string
		err      error
		testName string
	}{
		{
			note: &model.Note{
				Title: "test_title",
				Info:  "test_info",
			},
			tags: map[string]string{
				"1": "tag1",
				"2": "tag2",
			},
			expected: "",
			err:      nil,
			testName: "Test-1-OK",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.testName, func(t *testing.T) {
			db.SetUp()
			defer db.TearDown()

			if err := db.InsertTestNote(testCase.note, "1"); err != nil {
				log.Fatalln(err.Error())
			}

			storage := NewNoteStorage(db.Client)

			actual, err := storage.SetTags("1", testCase.tags)
			if err != nil {
				log.Fatalln(err.Error())
			}

			require.Equal(t, testCase.expected, actual)
			require.NoError(t, testCase.err, err)
		})
	}
}

func TestNoteStorage_RemoveTags(t *testing.T) {
	db := database.NewTestDBClient()
	defer db.Close()

	testTable := []struct {
		note     *model.Note
		tags     map[string]string
		expected string
		err      error
		testName string
	}{
		{
			note: &model.Note{
				Title: "test_title",
				Info:  "test_info",
			},
			tags: map[string]string{
				"1": "tag1",
				"2": "tag2",
			},
			expected: "",
			err:      nil,
			testName: "Test-1-OK",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.testName, func(t *testing.T) {
			db.SetUp()
			defer db.TearDown()

			if err := db.InsertTestNote(testCase.note, "1"); err != nil {
				log.Fatalln(err.Error())
			}

			storage := NewNoteStorage(db.Client)

			actual, err := storage.RemoveTags("1", testCase.tags)
			if err != nil {
				log.Fatalln(err.Error())
			}

			require.Equal(t, testCase.expected, actual)
			require.NoError(t, testCase.err, err)
		})
	}
}
