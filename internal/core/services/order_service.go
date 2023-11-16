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

func (os *OrderService) Buy(coinID int, amount float64) error {
	return nil
}

func (os *OrderService) Sell(coinID int, amount float64) error {
	return nil
}

func (os *OrderService) Cancel() error {
	return nil
}
