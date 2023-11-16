package services

import "Qexchange/internal/core/models"

type OrderDBService struct {
	//infrastructure.OrderDBRepository
}

func NewOrderDBService() *OrderDBService {
	return &OrderDBService{}
}

func (os *OrderDBService) GetUserBalance(int) (float64, error) {
	return 0, nil
}

func (os *OrderDBService) GetCoinPrice(int) (float64, error) {
	return 0, nil
}

func (os *OrderDBService) GetCoinCommission(int) (float64, error) {
	return 0, nil
}

func (os *OrderDBService) CommitOrder(models.Order) (int, error) {
	return 0, nil
}

func (os *OrderDBService) ChangeOrderStatus(int, int) error {
	return nil
}

func (os *OrderDBService) ValidateUser(int) error {
	return nil
}

func (os *OrderDBService) ValidateCoin(int) error {
	return nil
}

func (os *OrderDBService) ValidateOrder(int) error {
	return nil
}

func (os *OrderDBService) ValidateOrderBelongsToUser(int, int) error {
	return nil
}
