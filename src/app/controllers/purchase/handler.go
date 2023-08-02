package purchase

import (
	purchaseService "bank-app/src/app/controllers/purchase/services"
	"net/http"
)

type PurchaseHandler struct {
	s *purchaseService.PurchaseServiceI
}

func (p *PurchaseHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

}

func (p *PurchaseHandler) BuyGood(w http.ResponseWriter, r *http.Request) {

}
