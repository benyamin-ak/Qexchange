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
	balance, price, commission, err := os.validateData(o)
	if err != nil {
		return math.MinInt, err
	}
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

func (os *OrderService) validateData(o models.Order) (float64, float64, float64, error) {
	balance, err := os.dbc.GetUserBalance(o.UserID)
	if err != nil {
		return 0, 0, 0, err
	}
	price, err := os.dbc.GetCoinPrice(o.CoinID)
	if err != nil {
		return 0, 0, 0, err
	}
	commission, err := os.dbc.GetCoinCommission(o.CoinID)
	if err != nil {
		return 0, 0, 0, err
	}
	return balance, price, commission, nil
}
