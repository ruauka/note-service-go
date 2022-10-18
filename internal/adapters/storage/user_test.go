package storage

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"

	"web/internal/domain/entities/dto"
	"web/internal/domain/entities/model"
	"web/pkg/database"
)

func TestUserStorage_GetUserByID(t *testing.T) {
	db := database.NewTestDBClient()
	defer db.Close()

	testTable := []struct {
		user     *model.User
		userID   string
		expected *dto.UserResp
		err      error
		testName string
	}{
		{
			user: &model.User{
				Username: "test_name",
				Password: "test_password",
			},
			userID: "1",
			expected: &dto.UserResp{
				ID:       "1",
				Username: "test_name",
			},
			err:      nil,
			testName: "Test-1-OK",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.testName, func(t *testing.T) {
			db.SetUp()
			defer db.TearDown()

			if err := db.InsertTestUser(testCase.user); err != nil {
				log.Fatalln(err.Error())
			}

			storage := NewUserStorage(db.Client)

			actual, err := storage.GetUserByID(testCase.userID)
			if err != nil {
				log.Fatalln(err.Error())
			}

			require.Equal(t, testCase.expected, actual)
			require.NoError(t, testCase.err, err)
		})
	}
}

func TestUserStorage_GetAllUsers(t *testing.T) {
	db := database.NewTestDBClient()
	defer db.Close()

	testTable := []struct {
		user     *model.User
		expected []dto.UserResp
		err      error
		testName string
	}{
		{
			user: &model.User{
				Username: "test_name",
				Password: "test_password",
			},
			expected: []dto.UserResp{
				{
					ID:       "1",
					Username: "test_name",
				},
			},
			err:      nil,
			testName: "Test-1-OK",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.testName, func(t *testing.T) {
			db.SetUp()
			defer db.TearDown()

			if err := db.InsertTestUser(testCase.user); err != nil {
				log.Fatalln(err.Error())
			}

			storage := NewUserStorage(db.Client)

			actual, err := storage.GetAllUsers()
			if err != nil {
				log.Fatalln(err.Error())
			}

			require.Equal(t, testCase.expected, actual)
			require.NoError(t, testCase.err, err)
		})
	}
}

func TestUserStorage_UpdateUser(t *testing.T) {
	db := database.NewTestDBClient()
	defer db.Close()

	userName := "test_name1"
	userPassword := "test_password1"

	testTable := []struct {
		user     *model.User
		newUser  *dto.UserUpdate
		userID   string
		expected error
		testName string
	}{
		{
			user: &model.User{
				Username: "test_name",
				Password: "test_password",
			},
			newUser: &dto.UserUpdate{
				Username: &userName,
				Password: &userPassword,
			},
			userID:   "1",
			expected: nil,
			testName: "Test-1-OK",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.testName, func(t *testing.T) {
			db.SetUp()
			defer db.TearDown()

			if err := db.InsertTestUser(testCase.user); err != nil {
				log.Fatalln(err.Error())
			}

			storage := NewUserStorage(db.Client)

			actual := storage.UpdateUser(testCase.newUser, testCase.userID)

			require.NoError(t, testCase.expected, actual)
		})
	}
}

func TestUserStorage_DeleteUser(t *testing.T) {
	db := database.NewTestDBClient()
	defer db.Close()

	testTable := []struct {
		user     *model.User
		userID   string
		expected int
		err      error
		testName string
	}{
		{
			user: &model.User{
				Username: "test_name",
				Password: "test_password",
			},
			userID:   "1",
			expected: 1,
			err:      nil,
			testName: "Test-1-OK",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.testName, func(t *testing.T) {
			db.SetUp()
			defer db.TearDown()

			if err := db.InsertTestUser(testCase.user); err != nil {
				log.Fatalln(err.Error())
			}

			storage := NewUserStorage(db.Client)

			actual, err := storage.DeleteUser(testCase.userID)
			if err != nil {
				log.Fatalln(err.Error())
			}

			require.Equal(t, testCase.expected, actual)
			require.NoError(t, testCase.err, err)
		})
	}
}
