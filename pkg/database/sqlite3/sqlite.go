package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

func New(storagePath string) (*Storage, error) {
	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		fmt.Printf("Error of creating DB: %s", err)
		return nil, err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXIST users(
			id INTEGER PRIMARY KEY,
			name TEXT NOT NULL,
			email TEXT NOT NULL,
			password TEXT NOT NULL
		);
		CREATE TABLE IF NOT EXIST transactions(
			id INTEGER PRIMARY KEY,
			user_id INTEGER NOT NULL,
			amount DECIMAL(10, 2) NOT NULL,
			transaction_date TEXT NOT NULL,
			transaction_type TEXT NOT NULL,
			FOREIGN KEY (user_id) REFERENCES users(id)
		);
		CREATE TABLE IF NOT EXIST purchases(
			id INTEGER PRIMARY KEY,
			title TEXT NOT NULL,
			price DECIMAL(10, 2) NOT NULL,
			description TEXT NOT NULL
		);
	`)
	if err != nil {
		fmt.Printf("Error of creating tables: %s", err)
		return nil, err
	}

	return &Storage{db}, nil
}

//func (s *Storage) GetById(id int) (interface{}, error) {
//
//}
//
//func (s *Storage) GetAll() (interface{}, error) {
//
//}
//
//func (s *Storage) Insert(data interface{}) error {
//
//}
//
//func (s *Storage) Update(data interface{}) error {
//
//}
//
//func (s *Storage) Delete(id int) error {
//
//}

//func (storage *Storage) Get
