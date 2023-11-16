package services

import (
	"Qexchange/internal/core/contracts"
	"Qexchange/internal/core/models"
	"errors"
	"math"
	"time"
)

type OrderService struct {
	dbc contracts.OrderDBContract
}

func NewOrderService() *OrderService {
	return &OrderService{
		dbc: NewOrderDBService(),
	}
}

func (os *OrderService) Buy(userID int, coinID int, amount float64) (int, error) {
	o := models.Order{
		UserID:    userID,
		Side:      "buy",
		CoinID:    coinID,
		Quantity:  amount,
		Timestamp: time.Now(),
		Status:    models.OrderStatusActive,
	}
	if err := os.validateData(o); err != nil {
		return math.MinInt, err
	}
	balance := os.dbc.GetUserBalance(userID)
	price := os.dbc.GetCoinPrice(coinID)
	commission := os.dbc.GetCoinCommission(coinID)
	_ = commission
	if balance < amount*price {
		return math.MinInt, errors.New("insufficient funds")
	}
	ID, err := os.dbc.CommitOrder(o)
	if err != nil {
		return math.MinInt, err
	}
	return ID, nil
}

func (os *OrderService) Sell(userID int, coinID int, amount float64) (int, error) {
	return 0, nil
}

func (os *OrderService) Cancel(userID int, orderID int, userPassword string) error {
	return nil
}

func (os *OrderService) validateData(o models.Order) error {
	if err := os.dbc.ValidateUser(o.UserID); err != nil {
		return err
	}
	if err := os.dbc.ValidateCoin(o.CoinID); err != nil {
		return err
	}
	return nil
}
