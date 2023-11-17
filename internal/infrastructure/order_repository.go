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
	os.DB.Table("user").Where("id = ?", userID).Select("balance").Scan(&b)
	if b == -1.0 {
		return 0, errors.New("user not found")
	}
	return b, nil
}

func (os *OrderRepository) GetCoinPrice(coinID int) (float64, error) {
	p := -1.0
	os.DB.Table("mock_coin_price").Where("id = ?", coinID).Select("price").Scan(&p)
	if p == -1.0 {
		return 0, errors.New("coin not found")
	}
	return p, nil
}

func (os *OrderRepository) GetCoinCommission(coinID int) (float64, error) {
	c := -1.0
	os.DB.Table("commission").Where("coin_id = ?", coinID).Select("rate").Scan(&c)
	if c == -1.0 {
		return 0, errors.New("coin not found")
	}
	return c, nil
}

func (os *OrderRepository) CreateOrder(o *models.Order) (int, error) {
	ID := -1
	os.DB.Create(&o).Select("max(id)").Scan(&ID)
	if ID == -1 {
		return 0, errors.New("create order failed")
	}
	return ID, nil
}

func (os *OrderRepository) SubmitOrder(*models.Order) {

}

func (os *OrderRepository) ChangeOrderStatus(o *models.Order, status string) error {
	err := os.DB.Model(&models.Order{}).Where("id = ?", o.OrderID).Update("status", status)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (os *OrderRepository) ValidateUserPassword(int, string) error {
	return nil
}

func (os *OrderRepository) validateOrderBelongToUser(o *models.Order, userID int) error {
	order := &models.Order{}
	os.DB.Table("order").Where("id = ?", o.OrderID).Scan(&order)
	if order.UserID != userID {
		return errors.New("order does not belong to user")
	}
	return nil
}
