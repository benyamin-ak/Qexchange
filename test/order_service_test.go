package test

import (
	"Qexchange/internal/core/ordering"
	"Qexchange/internal/core/ordering/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

var os *ordering.OrderService = ordering.NewOrderService()

func TestBuyInsufficientFunds(t *testing.T) {
	o := &models.Order{
		UserID:   1,
		Side:     "buy",
		CoinID:   1,
		Quantity: 100000000,
	}
	_, err := os.Buy(o.UserID, o.CoinID, o.Quantity)
	assert.NotNil(t, err)
}

func TestBuy(t *testing.T) {
	for i := 0; i < 100; i++ {
		o := &models.Order{
			UserID:   i,
			Side:     "buy",
			CoinID:   Int(1, 10),
			Quantity: 0,
		}

		_, err := os.Buy(o.UserID, o.CoinID, o.Quantity)
		assert.Nil(t, err)
	}
}

func TestSell(t *testing.T) {
	for i := 0; i < 100; i++ {
		o := &models.Order{
			UserID:   i,
			Side:     "sell",
			CoinID:   Int(1, 10),
			Quantity: 0,
		}

		_, err := os.Sell(o.UserID, o.CoinID, o.Quantity)
		assert.Nil(t, err)
	}
}

func TestSellInsufficientFunds(t *testing.T) {
	o := &models.Order{
		UserID:   1,
		Side:     "sell",
		CoinID:   1,
		Quantity: 100000000,
	}
	_, err := os.Sell(o.UserID, o.CoinID, o.Quantity)
	assert.NotNil(t, err)
}
