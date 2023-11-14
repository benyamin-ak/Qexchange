package contracts

type OrderContract interface {
	//buy(coinID int, amount float64) error
	buy(int, float64) error
	//sell(coinID int, amount float64) error
	sell(int, float64) error
	cancel() error
}
