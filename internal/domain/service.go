package domain

type OrderService interface {
	PlaceOrder(order Order) error
	GetOrder(id string) (Order, error)
}
