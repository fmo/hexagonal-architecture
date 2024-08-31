package ports

import "github.com/fmo/hexagonal-architecture/internal/domain"

type OrderRepository interface {
	Save(order domain.Order) error
	FindByID(id string) (domain.Order, error)
}
