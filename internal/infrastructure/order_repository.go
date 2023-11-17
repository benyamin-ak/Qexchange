package infrastructure

import (
	"Qexchange/internal/core/models"

	"gorm.io/gorm"
)

type OrderRepository struct {
	gormDB *gorm.DB
}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{
		// gorm connection
	}
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

func (os *OrderRepository) ChangeOrderStatus(models.Order, string) error {
	return nil
}

func (os *OrderRepository) ValidateUserPassword(int, string) error {
	return nil
}
