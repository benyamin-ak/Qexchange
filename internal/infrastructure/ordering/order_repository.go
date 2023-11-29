package ordering

import (
	"Qexchange/internal/core/ordering/models"
	"Qexchange/internal/infrastructure"
	"errors"

	"gorm.io/gorm"
)

type OrderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{
		//DB = gorm connection
	}
}

func (os *OrderRepository) GetUserBalance(userID int, coinID int) (float64, error) {
	var b float64
	err := os.DB.Table(infrastructure.AssetTable).Where("user_id = ? AND coin_id = ?", userID, coinID).Select("quantity").Scan(&b)
	if err.Error != nil {
		return 0, err.Error
	}
	return b, nil
}

func (os *OrderRepository) GetCoinPrice(coinID int) (float64, error) {
	var p float64
	err := os.DB.Table(infrastructure.CoinPrices).Where("coin_id = ?", coinID).Select("price").Scan(&p)
	if err.Error != nil {
		return 0, err.Error
	}
	return p, nil
}

func (os *OrderRepository) GetCoinCommission(coinID int) (float64, error) {
	var c float64
	err := os.DB.Table(infrastructure.CommissionTable).Where("coin_id = ?", coinID).Select("rate").Scan(&c)
	if err.Error != nil {
		return 0, err.Error
	}
	return c, nil
}

func (os *OrderRepository) CreateOrder(o *models.Order) (int, error) {
	//ID ???
	err := os.DB.Table(infrastructure.OrderTable).Create(o)
	if err.Error != nil {
		return 0, err.Error
	}
	return o.OrderID, nil
}

func (os *OrderRepository) SubmitOrder(o *models.Order) {
	c, _ := os.GetCoinCommission(o.CoinID)
	balanceToDeduct := (1 + c) * o.Quantity * o.Price
	os.DB.Transaction(func(db *gorm.DB) error {
		if o.Side == "buy" {
			err := db.Table("assets").Where("user_id = ? AND coin_id = ?", o.UserID, 0 /*IRR coinID*/).Update("balance", gorm.Expr("balance - ?", balanceToDeduct))
			if err.Error != nil {
				return err.Error
			}
			err = db.Table("assets").Where("user_id = ? AND coin_id = ?", o.UserID, o.CoinID).Update("balance", gorm.Expr("balance + ?", o.Quantity))
			if err.Error != nil {
				return err.Error
			}
		} else {
			err := db.Table("assets").Where("user_id = ? AND coin_id = ?", o.UserID, o.CoinID).Update("balance", gorm.Expr("balance - ?", balanceToDeduct))
			if err.Error != nil {
				return err.Error
			}
			err = db.Table("assets").Where("user_id = ? AND coin_id = ?", o.UserID, 0 /*IRR coinID*/).Update("balance", gorm.Expr("balance + ?", o.Quantity*o.Price))
			if err.Error != nil {
				return err.Error
			}
		}
		return nil
	})
}

func (os *OrderRepository) ChangeOrderStatus(o *models.Order, status string) error {
	err := os.DB.Table(infrastructure.OrderTable).Where("id = ?", o.OrderID).Update("status", status)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (os *OrderRepository) ValidateOrderBelongToUser(o *models.Order, userID int) error {
	order := &models.Order{}
	os.DB.Table(infrastructure.OrderTable).Where("id = ?", o.OrderID).Scan(&order)
	if order.UserID != userID {
		return errors.New("order does not belong to user")
	}
	return nil
}
