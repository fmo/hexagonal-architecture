package domain

import (
	"github.com/fmo/hexagonal-architecture/internal/ports"
)

type OrderServiceImpl struct {
	repo ports.OrderRepository
}

func NewOrderService(repo ports.OrderRepository) *OrderServiceImpl {
	return &OrderServiceImpl{repo: repo}
}

func (s *OrderServiceImpl) PlaceOrder(order Order) error {
	return s.repo.Save(order)
}

func (s *OrderServiceImpl) GetOrder(id string) (Order, error) {
	return s.repo.FindByID(id)
}
