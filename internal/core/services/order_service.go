package services

import (
	"Qexchange/internal/core/contracts"
	"Qexchange/internal/core/models"
	"Qexchange/internal/infrastructure"
	"errors"
	"math"
	"time"
)

type OrderService struct {
	dbc contracts.OrderDBContract
}

func NewOrderService() *OrderService {
	return &OrderService{
		dbc: infrastructure.NewOrderRepository(),
	}
}

func (os *OrderService) Buy(userID int, coinID int, amount float64) (int, error) {
	o := &models.Order{
		UserID:    userID,
		Side:      "buy",
		CoinID:    coinID,
		Quantity:  amount,
		Timestamp: time.Now(),
		Status:    models.OrderStatusActive,
	}
	ID, err := os.dbc.CreateOrder(o)
	o.OrderID = ID
	if err != nil {
		os.dbc.ChangeOrderStatus(o, models.OrderStatusCancelled)
		return math.MinInt, err
	}
	balance, price, commission, err := os.validateData(o)
	if err != nil {
		os.dbc.ChangeOrderStatus(o, models.OrderStatusCancelled)
		return math.MinInt, err
	}
	o.Price = price
	_ = commission
	if balance < (1+commission)*amount*price {
		os.dbc.ChangeOrderStatus(o, models.OrderStatusCancelled)
		return math.MinInt, errors.New("insufficient funds")
	}
	return ID, nil
}

func (os *OrderService) Sell(userID int, coinID int, amount float64) (int, error) {
	o := &models.Order{
		UserID:    userID,
		Side:      "sell",
		CoinID:    coinID,
		Quantity:  amount,
		Timestamp: time.Now(),
		Status:    models.OrderStatusActive,
	}
	ID, err := os.dbc.CreateOrder(o)
	o.OrderID = ID
	balance, price, commission, err := os.validateData(o)
	if err != nil {
		os.dbc.ChangeOrderStatus(o, models.OrderStatusCancelled)
		return math.MinInt, err
	}
	o.Price = price
	if balance < (1+commission)*amount {
		os.dbc.ChangeOrderStatus(o, models.OrderStatusCancelled)
		return math.MinInt, errors.New("insufficient funds")
	}
	return ID, nil
}

func (os *OrderService) Cancel(userID int, orderID int, userPassword string) error {
	err := os.dbc.ValidateUserPassword(userID, userPassword)
	if err != nil {
		return err
	}
	o := &models.Order{
		UserID:  userID,
		OrderID: orderID,
	}
	err = os.dbc.ChangeOrderStatus(o, models.OrderStatusCancelled)
	if err != nil {
		return err
	}
	return nil
}

func (os *OrderService) validateData(o *models.Order) (float64, float64, float64, error) {
	var (
		balance    float64
		price      float64
		commission float64
		err        error
	)
	if o.Side == "buy" {
		balance, err = os.dbc.GetUserBalance(o.UserID, 0 /* IRR asset ID*/)
	} else {
		balance, err = os.dbc.GetUserBalance(o.UserID, o.CoinID)
	}
	if err != nil {
		return 0, 0, 0, err
	}
	price, err = os.dbc.GetCoinPrice(o.CoinID)
	if err != nil {
		return 0, 0, 0, err
	}
	commission, err = os.dbc.GetCoinCommission(o.CoinID)
	if err != nil {
		return 0, 0, 0, err
	}
	return balance, price, commission, nil
}
