package db

import (
	"database/sql"
	"github.com/fmo/hexagonal-architecture/internal/domain"
	"github.com/fmo/hexagonal-architecture/internal/ports"
)

type SqlOrderRepository struct {
	db *sql.DB
}

func NewSqlOrderRepository(db *sql.DB) ports.OrderRepository {
	return &SqlOrderRepository{db: db}
}

func (r *SqlOrderRepository) Save(order domain.Order) error {
	_, err := r.db.Exec("INSERT INTO orders (id,, amount, status) VALUES (?, ?, ?)", order.ID, order.Amount, order.Status)

	return err
}

func (r *SqlOrderRepository) FindByID(id string) (domain.Order, error) {
	row := r.db.QueryRow("SELECT id, amount, status FROM orders WHERE id = ?", id)
	var order domain.Order
	err := row.Scan(&order.ID, &order.Amount, &order.Status)
	if err != nil {
		return domain.Order{}, err
	}

	return order, nil
}
