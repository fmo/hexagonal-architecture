package ports

import "github.com/fmo/hexagonal-architecture/internal/domain"

type OrderService interface {
	PlaceOrder(order domain.Order) error
	GetOrder(id string) (domain.Order, error)
}
