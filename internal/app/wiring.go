package app

import (
	"database/sql"
	"github.com/fmo/hexagonal-architecture/internal/adapters/inbound/controller"
	"github.com/fmo/hexagonal-architecture/internal/adapters/outbound/db"
	"github.com/fmo/hexagonal-architecture/internal/domain/srv"
	"net/http"
)

func SetupApplication(database *sql.DB) http.Handler {
	repository := db.NewSqlOrderRepository(database)

	service := srv.NewOrderService(repository)

	orderController := controller.NewOrderController(service)

	mux := http.NewServeMux()
	mux.HandleFunc("/place-order", orderController.HandlePlaceOrder)

	return mux
}
