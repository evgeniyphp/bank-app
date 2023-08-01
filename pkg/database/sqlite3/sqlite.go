package sqlite3

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	DB *sql.DB
}

func New(storagePath string) (*Storage, error) {
	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		fmt.Printf("Error of creating DB: %s\n", err)
		return nil, err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			email TEXT NOT NULL,
			password TEXT NOT NULL,
			amount DECIMAL(10, 2) DEFAULT 0
		);
		CREATE TABLE IF NOT EXISTS transactions(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER NOT NULL,
			amount DECIMAL(10, 2) NOT NULL,
			transaction_date TEXT NOT NULL,
			transaction_type TEXT NOT NULL,
			FOREIGN KEY (user_id) REFERENCES users(id)
		);
		CREATE TABLE IF NOT EXISTS purchases(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			price DECIMAL(10, 2) NOT NULL,
			description TEXT NOT NULL
		);
	`)
	if err != nil {
		fmt.Printf("Error of creating tables: %s\n", err)
		return nil, err
	}

	return &Storage{db}, nil
}
