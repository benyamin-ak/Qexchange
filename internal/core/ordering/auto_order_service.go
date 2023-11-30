package ordering

import (
	"Qexchange/internal/core/ordering/models"
	"Qexchange/internal/infrastructure/ordering"
	"time"
)

type AutoOrderService struct {
	occ           OrderCoreContract
	odc           OrderDBContract
	order         *models.Order
	PriceToSubmit float64
	Done          bool
}

func NewAutoOrderService() *AutoOrderService {
	return &AutoOrderService{
		occ: NewOrderService(),
		odc: ordering.NewOrderRepository(),
	}
}

func (aos *AutoOrderService) StartPolling(UserID int, CoinID int, Quantity float64, Side string, pts float64) {
	aos.order = &models.Order{
		UserID:   UserID,
		CoinID:   CoinID,
		Quantity: Quantity,
		Side:     Side,
		Status:   models.OrderStatusActive,
	}
	go aos.CheckOrders()
}

func (aos *AutoOrderService) CheckOrders() {
	for {
		p, err := aos.odc.GetCoinPrice(aos.order.CoinID)
		if err != nil {
			aos.Done = true
			break
		}
		if aos.order.Side == "buy" {
			if p <= aos.order.Price {
				aos.Done = true
				_, err := aos.occ.Buy(aos.order.UserID, aos.order.CoinID, aos.order.Quantity)
				if err != nil {
					break
				}
			}
		} else {
			if p >= aos.PriceToSubmit {
				aos.Done = true
				_, err := aos.occ.Sell(aos.order.UserID, aos.order.CoinID, aos.order.Quantity)
				if err != nil {
					break
				}
			}
		}
		time.Sleep(10 * time.Second)
	}
}
