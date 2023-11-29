package models

type Order struct {
	OrderID  int     `gorm:"column:id"`
	UserID   int     `gorm:"column:user_id"`
	Side     string  `gorm:"column:side"`
	Quantity float64 `gorm:"column:quantity"`
	Price    float64 `gorm:"column:price"`
	CoinID   int     `gorm:"column:coin_id"`
	Status   string  `gorm:"column:status"`
}

const (
	OrderStatusActive    = "active"
	OrderStatusCancelled = "cancelled"
	OrderStatusCompleted = "completed"
)
