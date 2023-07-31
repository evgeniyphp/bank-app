package database

type Database interface {
	GetById(int) (interface{}, error)
	GetAll() (interface{}, error)
	Insert(interface{}) error
	Update(interface{}) error
	Delete(int) error
}
