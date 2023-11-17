package services

import "Qexchange/internal/core/models"

type OrderDBService struct {
	//infrastructure.OrderDBRepository
}

func NewOrderRepository() *OrderDBService {
	return &OrderDBService{}
}

func (os *OrderDBService) GetUserBalance(int, int) (float64, error) {
	return 0, nil
}

func (os *OrderDBService) GetCoinPrice(int) (float64, error) {
	return 0, nil
}

func (os *OrderDBService) GetCoinCommission(int) (float64, error) {
	return 0, nil
}

func (os *OrderDBService) CreateOrder(models.Order) (int, error) {
	return 0, nil
}

func (os *OrderDBService) SubmitOrder(models.Order) {
	return
}

func (os *OrderDBService) ChangeOrderStatus(models.Order, int) error {
	return nil
}

func (os *OrderDBService) ValidateUserPassword(int, string) error {
	return nil
}
