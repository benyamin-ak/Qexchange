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

func NewAutoOrderService(UserID int, CoinID int, Quantity float64, Side string, pts float64) *AutoOrderService {
	return &AutoOrderService{
		occ: NewOrderService(),
		odc: ordering.NewOrderRepository(),
		order: &models.Order{
			UserID:   UserID,
			Side:     Side,
			CoinID:   CoinID,
			Quantity: Quantity,
			Status:   models.OrderStatusActive,
		},
		PriceToSubmit: pts,
		Done:          false,
	}
}

func (aos *AutoOrderService) StartPolling() {
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
