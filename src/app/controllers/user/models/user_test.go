package models

import (
	"bank-app/pkg/database/sqlite3"
	"testing"
)

func TestMain(t *testing.M) {
	t.Run()
}

func TestGetById(t *testing.T) {
	storage, err := sqlite3.New(":memory:")
	db := storage.DB

	defer db.Close()

	if err != nil {
		t.Fatal("Cannot create db for tests")
	}

	_, err = db.Exec(`
		INSERT INTO users (name, email, password, amount) 
		VALUES (?, ?, ?, ?)`,
		"test1", "test1@m", "123", 100.00,
	)
	if err != nil {
		t.Fatalf("Failed executing insert %s", err)
	}

	repo := NewUserRepository(db)

	user, err := repo.GetById(1)

	expectedUser := &User{
		ID:       1,
		Name:     "test1",
		Email:    "test1@m",
		Password: "123",
		Amount:   100.00,
	}

	if user.Name != expectedUser.Name {
		t.Errorf("Expected name to be %s, but got %s", expectedUser.Name, user.Name)
	}
	if user.Email != expectedUser.Email {
		t.Errorf("Expected email to be %s, but got %s", expectedUser.Email, user.Email)
	}
	if user.Password != expectedUser.Password {
		t.Errorf("Expected password to be %s, but got %s", expectedUser.Password, user.Password)
	}
	if user.Amount != expectedUser.Amount {
		t.Errorf("Expected amoun to be %f, but got %f", expectedUser.Amount, user.Amount)
	}

}
