package test

import (
	"Qexchange/internal/core/ordering"
	"testing"
)

func TestBuyInsufficientFunds(t *testing.T) {
	orderService := ordering.NewOrderService()
	_ = orderService
}
