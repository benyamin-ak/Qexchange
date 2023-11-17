package infrastructure

import (
	"Qexchange/internal/core/models"
	"errors"

	"gorm.io/gorm"
)

type OrderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{
		// gorm connection
	}
}

func (os *OrderRepository) GetUserBalance(userID int, coinID int) (float64, error) {
	b := -1.0
	os.DB.Model(&models.User{}).Where("id = ?", userID).Select("balance").Scan(&b)
	if b == -1.0 {
		return 0, errors.New("user not found")
	}
	return b, nil
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
