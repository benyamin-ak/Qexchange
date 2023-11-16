package contracts

type OrderCoreContract interface {
	//buy(userID int, coinID int, amount float64) orderID error
	Buy(int, int, float64) (int, error)
	//sell(userID int, coinID int, amount float64) orderID error
	Sell(int, int, float64) (int, error)
	//cancel(userID int, orderID int) error
	Cancel(int, int) error
}

type OrderDBContract interface {
	ValidateUser(int) error
	ValidateCoin(int) error
	ValidateOrder(int) error
	ValidateOrderBelongsToUser(int, int) error
	GetUserBalance(int) (float64, error)
	GetCoinPrice(int) (float64, error)
	GetCoinCommission(int) (float64, error)
	CommitOrder(int, int, string, float64, float64, int) (int, error)
	ChangeOrderStatus(int, string) error
}
