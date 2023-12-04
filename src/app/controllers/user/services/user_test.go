package services

import (
	"bank-app/pkg/database/sqlite3"
	"bank-app/src/app/controllers/user/models"
	"fmt"
	"testing"
)

var storage *sqlite3.Storage

func TestMain(t *testing.M) {
	// storage, err := sqlite3.New(":memory:")

	// db := storage.DB
	// if err != nil {
	// 	fmt.Print("Cannot create db for tests")
	// }
	// defer db.Close()
	t.Run()
}

func TestCreateUser(t *testing.T) {
	storage, err := sqlite3.New(":memory:")

	db := storage.DB
	if err != nil {
		fmt.Print("Cannot create db for tests")
	}
	defer db.Close()

	userRepository := models.NewUserRepository(db)

	testTable := []struct {
		input    models.User
		expected any
	}{
		{
			input: models.User{
				Name:     "test",
				Email:    "test@m",
				Password: "123",
			},
			expected: nil,
		},
		{
			input: models.User{
				Name:     "",
				Password: "123",
			},
			expected: nil,
		},
	}

	for i, testCase := range testTable {
		t.Run("Ok", func(t *testing.T) {
			t.Log(i)
			err := userRepository.Insert(&testCase.input)
			t.Log(err)
			if err != testCase.expected {
				t.Errorf("Failed test: expected %s but was got %s", testCase.expected, err)
			}
		})
	}

}
