package services

import "Qexchange/internal/core/models"

type OrderRepository struct {
	//infrastructure.OrderDBRepository
}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{}
}

func (os *OrderRepository) GetUserBalance(int, int) (float64, error) {
	return 0, nil
}

func (os *OrderRepository) GetCoinPrice(int) (float64, error) {
	return 0, nil
}

func (os *OrderRepository) GetCoinCommission(int) (float64, error) {
	return 0, nil
}

func (os *OrderRepository) CreateOrder(models.Order) (int, error) {
	return 0, nil
}

func (os *OrderRepository) SubmitOrder(models.Order) {
	return
}

func (os *OrderRepository) ChangeOrderStatus(models.Order, int) error {
	return nil
}

func (os *OrderRepository) ValidateUserPassword(int, string) error {
	return nil
}
