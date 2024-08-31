package srv

import (
	"github.com/fmo/hexagonal-architecture/internal/domain"
	"github.com/fmo/hexagonal-architecture/internal/ports"
)

type OrderServiceImpl struct {
	repo ports.OrderRepository
}

func NewOrderService(repo ports.OrderRepository) *OrderServiceImpl {
	return &OrderServiceImpl{repo: repo}
}

func (s *OrderServiceImpl) PlaceOrder(order domain.Order) error {
	return s.repo.Save(order)
}

func (s *OrderServiceImpl) GetOrder(id string) (domain.Order, error) {
	return s.repo.FindByID(id)
}
