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
	ValidateUser(int) error
	ValidateCoin(int) error
	ValidateOrder(int) error
	ValidateOrderBelongsToUser(int, int) error
	GetUserBalance(int) (float64, error)
	GetCoinPrice(int) (float64, error)
	GetCoinCommission(int) (float64, error)
	CommitOrder(models.Order) (int, error)
	ChangeOrderStatus(int, int) error
}
