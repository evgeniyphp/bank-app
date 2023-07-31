package user_model

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

func (s *User) GetById(id int) (interface{}, error) {

}

func (s *User) GetAll() (interface{}, error) {

}

func (s *User) Insert(data interface{}) error {

}

func (s *User) Update(data interface{}) error {

}

func (s *User) Delete(id int) error {

}
