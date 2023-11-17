package contracts

import "Qexchange/internal/core/models"

type OrderCoreContract interface {
	//buy(userID int, coinID int, amount float64) orderID error
	Buy(int, int, float64) (int, error)
	//sell(userID int, coinID int, amount float64) orderID error
	Sell(int, int, float64) (int, error)
	//cancel(userID int, orderID int, userPassword string) error
	Cancel(int, int, string) error
}

type OrderDBContract interface {
	GetUserBalance(int, int) (float64, error)
	GetCoinPrice(int) (float64, error)
	GetCoinCommission(int) (float64, error)
	CreateOrder(models.Order) (int, error)
	SubmitOrder(models.Order) (int, error)
	ChangeOrderStatus(models.Order, int) error
	ValidateUserPassword(int, string) error
}
