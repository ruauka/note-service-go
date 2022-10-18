package storage

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"

	"web/internal/domain/entities/dto"
	"web/internal/domain/entities/model"
	"web/pkg/database"
)

func TestTagStorage_CreateTag(t *testing.T) {
	db := database.NewTestDBClient()
	defer db.Close()

	testTable := []struct {
		tag, expected *model.Tag
		err           error
		testName      string
	}{
		{
			tag: &model.Tag{
				TagName: "test_tag",
			},
			expected: &model.Tag{
				ID:      "1",
				TagName: "test_tag",
			},
			err:      nil,
			testName: "Test-1-OK",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.testName, func(t *testing.T) {
			db.SetUp()
			defer db.TearDown()

			storage := NewTagStorage(db.Client)

			actual, err := storage.CreateTag(testCase.tag, "1")
			if err != nil {
				log.Fatalln(err.Error())
			}

			require.Equal(t, testCase.expected, actual)
			require.NoError(t, testCase.err, err)
		})
	}
}

func TestTagStorage_GetTagByID(t *testing.T) {
	db := database.NewTestDBClient()
	defer db.Close()

	testTable := []struct {
		tag      *model.Tag
		expected *dto.TagResp
		err      error
		testName string
	}{
		{
			tag: &model.Tag{
				TagName: "test_tag",
			},
			expected: &dto.TagResp{
				TagName: "test_tag",
			},
			testName: "Test-1-OK",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.testName, func(t *testing.T) {
			db.SetUp()
			defer db.TearDown()

			if err := db.InsertTestTag(testCase.tag, "1"); err != nil {
				log.Fatalln(err.Error())
			}

			storage := NewTagStorage(db.Client)

			actual, err := storage.GetTagByID("1", "1")
			if err != nil {
				log.Fatalln(err.Error())
			}

			require.Equal(t, testCase.expected, actual)
			require.NoError(t, testCase.err, err)
		})
	}
}

func TestTagStorage_GetAllTagsByUser(t *testing.T) {
	db := database.NewTestDBClient()
	defer db.Close()

	testTable := []struct {
		tag      *model.Tag
		expected []dto.TagsResp
		err      error
		testName string
	}{
		{
			tag: &model.Tag{
				TagName: "test_tag",
			},
			expected: []dto.TagsResp{
				{
					ID:      "1",
					TagName: "test_tag",
				},
			},
			testName: "Test-1-OK",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.testName, func(t *testing.T) {
			db.SetUp()
			defer db.TearDown()

			if err := db.InsertTestTag(testCase.tag, "1"); err != nil {
				log.Fatalln(err.Error())
			}

			storage := NewTagStorage(db.Client)

			actual, err := storage.GetAllTagsByUser("1")
			if err != nil {
				log.Fatalln(err.Error())
			}

			require.Equal(t, testCase.expected, actual)
			require.NoError(t, testCase.err, err)
		})
	}
}

func TestTagStorage_UpdateTag(t *testing.T) {
	db := database.NewTestDBClient()
	defer db.Close()

	tagName := "new_tag1"

	testTable := []struct {
		tag      *model.Tag
		newTag   *dto.TagUpdate
		expected error
		testName string
	}{
		{
			tag: &model.Tag{
				TagName: "test_tag",
			},
			newTag: &dto.TagUpdate{
				TagName: &tagName,
			},
			expected: nil,
			testName: "Test-1-OK",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.testName, func(t *testing.T) {
			db.SetUp()
			defer db.TearDown()

			if err := db.InsertTestTag(testCase.tag, "1"); err != nil {
				log.Fatalln(err.Error())
			}

			storage := NewTagStorage(db.Client)

			actual := storage.UpdateTag(testCase.newTag, "1")

			require.NoError(t, testCase.expected, actual)
		})
	}
}

func TestTagStorage_DeleteTag(t *testing.T) {
	db := database.NewTestDBClient()
	defer db.Close()

	testTable := []struct {
		tag      *model.Tag
		expected int
		err      error
		testName string
	}{
		{
			tag: &model.Tag{
				TagName: "test_tag",
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
			
			if err := db.InsertTestTag(testCase.tag, "1"); err != nil {
				log.Fatalln(err.Error())
			}

			storage := NewTagStorage(db.Client)

			actual, err := storage.DeleteTag("1", "1")
			if err != nil {
				log.Fatalln(err.Error())
			}

			require.Equal(t, testCase.expected, actual)
			require.NoError(t, testCase.err, err)
		})
	}
}
