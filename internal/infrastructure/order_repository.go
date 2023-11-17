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

func (os *OrderRepository) GetCoinPrice(coinID int) (float64, error) {
	p := -1.0
	os.DB.Model(&models.MockCoinPrice{}).Where("id = ?", coinID).Select("price").Scan(&p)
	if p == -1.0 {
		return 0, errors.New("coin not found")
	}
	return p, nil
}

func (os *OrderRepository) GetCoinCommission(coinID int) (float64, error) {
	c := -1.0
	os.DB.Model(&models.Commission{}).Where("coin_id = ?", coinID).Select("rate").Scan(&c)
	if c == -1.0 {
		return 0, errors.New("coin not found")
	}
	return c, nil
}

func (os *OrderRepository) CreateOrder(o models.Order) (int, error) {
	ID := -1
	os.DB.Create(&o).Select("max(id)").Scan(&ID)
	if ID == -1 {
		return 0, errors.New("create order failed")
	}
	return ID, nil
}

func (os *OrderRepository) SubmitOrder(models.Order) {

}

func (os *OrderRepository) ChangeOrderStatus(models.Order, string) error {
	return nil
}

func (os *OrderRepository) ValidateUserPassword(int, string) error {
	return nil
}

func validateOrderBelongToUser() error {
	return nil
}
