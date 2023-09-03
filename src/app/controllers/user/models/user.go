package models

import (
	"database/sql"
	"fmt"
)

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
	Amount   float64
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (s *UserRepository) GetById(id int) (*User, error) {
	stmt, err := s.db.Prepare(`SELECT * FROM users WHERE id=?`)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return nil, err
	}

	var user User

	row := stmt.QueryRow(id)
	err = row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Amount)
	if err != nil {
		fmt.Printf("Error2: %s\n", err)
		return nil, err
	}

	return &user, nil
}

func (s *UserRepository) Insert(user *User) error {
	stmt, err := s.db.Prepare("INSERT INTO users(name, email, password) VALUES(?, ?, ?)")
	if err != nil {
		return err
	}

	result, err := stmt.Exec(user.Name, user.Email, user.Password)
	if err != nil {
		fmt.Printf("Error3: %s", err)
		return err
	}

	fmt.Println("Result of inserting: ", result)
	return nil
}

func (s *UserRepository) Update(user *User) error {
	stmt, err := s.db.Prepare("UPDATE users SET amount=? WHERE id=?")
	if err != nil {
		return err
	}

	result, err := stmt.Exec(user.Amount, user.ID)
	if err != nil {
		fmt.Printf("Error4: %s", err)
		return err
	}

	fmt.Println("Result of updating: ", result)
	return nil
}
