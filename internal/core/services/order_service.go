package services

import "Qexchange/internal/core/contracts"

type OrderService struct {
	dbc contracts.OrderDBContract
}

func NewOrderService() *OrderService {
	return &OrderService{
		dbc: NewOrderDBService(),
	}
}

func (os *OrderService) Buy(userID int, coinID int, amount float64) (int, error) {

}

func (os *OrderService) Sell(userID int, coinID int, amount float64) (int, error) {
	return 0, nil
}

func (os *OrderService) Cancel(userID int, orderID int) error {
	return nil
}
