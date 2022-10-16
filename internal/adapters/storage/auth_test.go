package storage

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"web/internal/domain/entities/model"
	"web/pkg/database/sqlite"
)

func TestUserAuthStorage_RegisterUser(t *testing.T) {
	db := sqlite.NewDBTestConn()
	defer db.Close()

	up := sqlite.FileOpen("../../../migrate/000001_init.up.sql")
	down := sqlite.FileOpen("../../../migrate/000001_init.down.sql")

	testTable := []struct {
		user, expected *model.User
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
			sqlite.SetUp(db, up)
			defer sqlite.TearDown(db, down)

			store := NewAuthStorage(db)

			actual, err := store.RegisterUser(testCase.user)
			if err != nil {
				fmt.Println(err.Error())
			}

			require.Equal(t, testCase.expected, actual)
		})
	}
}
