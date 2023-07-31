package transaction_model

type Transaction struct {
	ID              int
	UserID         int
	Amount          float64
	TransactionDate string
	TransactionType string
}

func (s *Transaction) GetById(id int) (interface{}, error) {

}

func (s *Transaction) GetAll() (interface{}, error) {

}

func (s *Transaction) Insert(data interface{}) error {

}

func (s *Transaction) Update(data interface{}) error {

}

func (s *Transaction) Delete(id int) error {

}