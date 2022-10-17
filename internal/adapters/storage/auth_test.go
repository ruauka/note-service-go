package storage

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"

	"web/internal/domain/entities/model"
	"web/pkg/database"
)

func TestUserAuthStorage_RegisterUser(t *testing.T) {
	db := database.NewTestDBClient()
	defer db.Close()

	testTable := []struct {
		user, expected *model.User
		err            error
		testName       string
	}{
		{
			user: &model.User{
				Username: "test_name",
				Password: "test_password",
			},
			expected: &model.User{
				ID:       "1",
				Username: "test_name",
				Password: "test_password",
			},
			testName: "Test-1-OK",
		},
		{
			user: &model.User{
				Username: "test_name1",
				Password: "test_password1",
			},
			expected: &model.User{
				ID:       "1",
				Username: "test_name1",
				Password: "test_password1",
			},
			testName: "Test-2-OK",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.testName, func(t *testing.T) {
			db.SetUp()
			defer db.TearDown()

			storage := NewAuthStorage(db.Client)

			actual, err := storage.RegisterUser(testCase.user)
			if err != nil {
				log.Fatalln(err.Error())
			}

			require.Equal(t, testCase.expected, actual)
			require.NoError(t, testCase.err, err)
		})
	}
}

func TestUserAuthStorage_GetUserForToken(t *testing.T) {
	db := database.NewTestDBClient()
	defer db.Close()

	testTable := []struct {
		user     *model.User
		expected *model.User
		err      error
		testName string
	}{
		{
			user: &model.User{
				Username: "test_name",
				Password: "test_password_hash",
			},
			expected: &model.User{
				ID: "1",
			},
			testName: "Test-1-OK",
		},
		{
			user: &model.User{
				Username: "test_name1",
				Password: "test_password_hash1",
			},
			expected: &model.User{
				ID: "1",
			},
			testName: "Test-2-OK",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.testName, func(t *testing.T) {
			db.SetUp()
			defer db.TearDown()

			storage := NewAuthStorage(db.Client)

			if err := db.UserInsert(testCase.user); err != nil {
				log.Fatalln(err.Error())
			}

			actual, err := storage.GetUserForToken(testCase.user.Username, testCase.user.Password)
			if err != nil {
				log.Fatalln(err.Error())
			}

			require.Equal(t, testCase.expected, actual)
			require.NoError(t, testCase.err, err)
		})
	}
}
