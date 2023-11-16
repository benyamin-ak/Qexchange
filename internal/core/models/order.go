package models

import "time"

type Order struct {
	orderID   int
	userID    int
	side      string
	quntity   float64
	price     float64
	coinID    int
	timestamp time.Time
	status    string
}
