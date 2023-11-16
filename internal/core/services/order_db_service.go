package services

type OrderDBService struct {
	//infrastructure.OrderDBRepository
}

func NewOrderDBService() *OrderDBService {
	return &OrderDBService{}
}

func (os *OrderDBService) GetUserBalance(int) (float64, error) {
	return 0, nil
}

func (os *OrderDBService) GetCoinPrice(int) (float64, error) {
	return 0, nil
}

func (os *OrderDBService) GetCoinCommission(int) (float64, error) {
	return 0, nil
}

func (os *OrderDBService) CommitOrder(int, int, string, float64, float64, int) (int, error) {
	return 0, nil
}

func (os *OrderDBService) ChangeOrderStatus(int, string) error {
	return nil
}
