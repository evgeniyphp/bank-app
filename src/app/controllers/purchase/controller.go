package purchase

import (
	purchaseModel "bank-app/src/app/controllers/purchase/models"
	purchaseService "bank-app/src/app/controllers/purchase/services"
	"encoding/json"
	"fmt"
	"net/http"
)

type Controller struct {
	s purchaseService.PurchaseServiceI
}

func NewPurchaseController(s purchaseService.PurchaseServiceI) *Controller {
	return &Controller{s}
}

func (p *Controller) CreateGood(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var data struct {
		Title       string  `json:"title"`
		Price       float64 `json:"price"`
		Description string  `json:"description"`
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println("Bad request", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	good := purchaseModel.Good{
		Title:       data.Title,
		Price:       data.Price,
		Description: data.Description,
	}

	err = p.s.CreateGood(&good)
	if err != nil {
		fmt.Println("Cannot create good", err)
		http.Error(w, "Cannot create good", http.StatusBadRequest)
		return
	}

	result, _ := json.Marshal(good)

	w.WriteHeader(http.StatusCreated)
	w.Write(result)
}

func (p *Controller) BuyGood(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var data struct {
		UserID int `json:"user_id"`
		GoodID int `json:"good_id"`
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println("Bad request", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	err = p.s.BuyGood(data.UserID, data.GoodID)
	if err != nil {
		fmt.Println("Cannot buy good", err)
		http.Error(w, "Cannot buy good", http.StatusBadRequest)
		return
	}
}
