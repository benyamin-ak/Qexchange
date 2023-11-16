package contracts

type OrderCoreContract interface {
	//buy(coinID int, amount float64) orderID error
	Buy(int, float64) (int, error)
	//sell(coinID int, amount float64) orderID error
	Sell(int, float64) (int, error)
	//cancel(orderID int) error
	Cancel(int) error
}

type OrderDBContract interface {
	GetUserBalance(int) (float64, error)
	GetCoinPrice(int) (float64, error)
	GetCoinCommission(int) (float64, error)
	CommitOrder(int, int, string, float64, float64, int) (int, error)
	ChangeOrderStatus(int, string) error
}
