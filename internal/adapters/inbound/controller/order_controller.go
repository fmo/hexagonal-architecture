package controller

import (
	"encoding/json"
	"github.com/fmo/hexagonal-architecture/internal/domain"
	"net/http"
)

type OrderController struct {
	service domain.OrderService
}

func NewOrderController(service domain.OrderService) *OrderController {
	return &OrderController{service: service}
}

func (c *OrderController) HandlePlaceOrder(w http.ResponseWriter, r *http.Request) {
	var order domain.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
