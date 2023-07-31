package purchase_model

type Good struct {
	ID          int
	Title       string
	Price       float64
	Description string
}

func (s *Good) GetById(id int) (interface{}, error) {

}

func (s *Good) GetAll() (interface{}, error) {

}

func (s *Good) Insert(data interface{}) error {

}

func (s *Good) Update(data interface{}) error {

}

func (s *Good) Delete(id int) error {

}
