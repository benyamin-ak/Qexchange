package ordering

type OrderRequest struct {
	// Will change to SessionID
	UserID int     `json:"user_id"`
	CoinID int     `json:"coin_id"`
	Amount float64 `json:"amount"`
}

type CancelRequest struct {
	// Will change to SessionID
	UserID  int `json:"user_id"`
	OrderID int `json:"order_id"`
}

type OrderResponse struct {
	OrderID int    `json:"order_id"`
	Error   string `json:"error"`
}
