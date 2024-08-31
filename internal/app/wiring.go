package app

import (
	"database/sql"
	"github.com/fmo/hexagonal-architecture/internal/adapters/inbound/controller"
	"github.com/fmo/hexagonal-architecture/internal/adapters/outbound/db"
	"github.com/fmo/hexagonal-architecture/internal/domain"
	"github.com/fmo/hexagonal-architecture/internal/ports"
	"net/http"
)

func SetupApplication(database *sql.DB) http.Handler {
	repository := db.NewSqlOrderRepository(database)

	service := NewOrderService(repository)

	orderController := controller.NewOrderController(service)

	mux := http.NewServeMux()
	mux.HandleFunc("/place-order", orderController.HandlePlaceOrder)

	return mux
}

type OrderService struct {
	repo ports.OrderRepository
}

func NewOrderService(repo ports.OrderRepository) domain.OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) PlaceOrder(order domain.Order) error {
	return s.repo.Save(order)
}

func (s *OrderService) GetOrder(id string) (domain.Order, error) {
	return s.repo.FindByID(id)
}
