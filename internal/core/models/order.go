package models

import "time"

type Order struct {
	OrderID   int
	UserID    int
	Side      string
	Quantity  float64
	Price     float64
	CoinID    int
	Timestamp time.Time
	Status    int
}

const (
	OrderStatusActive = iota
	OrderStatusCancelled
	OrderStatusCompleted
)
