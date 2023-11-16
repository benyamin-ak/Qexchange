package contracts

type OrderCoreContract interface {
	//buy(coinID int, amount float64) error
	Buy(int, float64) error
	//sell(coinID int, amount float64) error
	Sell(int, float64) error
	Cancel() error
}

type OrderDBContract interface {
	GetUserBalance(int) (float64, error)
	GetCoinPrice(int) (float64, error)
	GetCoinCommission(int) (float64, error)
	CommitOrder(int, int, string, float64, float64, int) error
	ChangeOrderStatus(int, string) error
}
