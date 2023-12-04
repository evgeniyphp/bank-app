// Package models describes the Good model and implements repository for that model

package models

import (
	"database/sql"
	"fmt"
)

type Good struct {
	ID          int
	Title       string
	Price       float64
	Description string
}

type GoodRepository struct {
	db *sql.DB
}

func NewGoodRepository(db *sql.DB) *GoodRepository {
	return &GoodRepository{db}
}

// Gets the good model by id
// Returns the good model and error
func (s *GoodRepository) GetById(id int) (*Good, error) {
	stmt, err := s.db.Prepare("SELECT * FROM purchases WHERE id=?")
	if err != nil {
		return nil, err
	}

	var good Good

	err = stmt.QueryRow(id).Scan(&good.ID, &good.Title, &good.Price, &good.Description)
	if err != nil {
		return nil, err
	}

	return &good, nil
}

func (s *GoodRepository) Insert(obj *Good) error {
	stmt, err := s.db.Prepare("INSERT INTO purchases(title, price, description) VALUES(?,?,?)")
	if err != nil {
		return err
	}
	r, err := stmt.Exec(obj.Title, obj.Price, obj.Description)
	if err != nil {
		return err
	}
	fmt.Println("Result insert good: ", r)

	return nil
}
