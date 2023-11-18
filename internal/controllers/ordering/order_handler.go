package ordering

import (
	"Qexchange/internal/core/ordering"
	"errors"
	"fmt"
	"math"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type OrderRequest struct {
	UserID int     `json:"user_id"`
	CoinID int     `json:"coin_id"`
	Amount float64 `json:"amount"`
}

type CancelRequest struct {
	UserID  int `json:"user_id"`
	OrderID int `json:"order_id"`
}

type Response struct {
	OrderID int    `json:"order_id"`
	Error   string `json:"error"`
}

func NewResponse(orderID int, err error) Response {
	if err == nil {
		err = fmt.Errorf("")
	}
	return Response{
		OrderID: orderID,
		Error:   err.Error(),
	}
}

type OrderHandler struct {
	osc ordering.OrderCoreContract
}

func NewOrderHandler() *OrderHandler {
	return &OrderHandler{
		osc: ordering.NewOrderService(),
	}
}

func (oh *OrderHandler) Buy(c echo.Context) error {
	or := new(OrderRequest)
	if err := c.Bind(or); err != nil {
		return c.JSON(http.StatusBadRequest, NewResponse(math.MinInt, err))
	}
	if or.Amount <= 0 || or.UserID <= 0 || or.CoinID <= 0 {
		return c.JSON(http.StatusBadRequest, NewResponse(math.MinInt, errors.New("invalid request")))
	}
	orderID, err := oh.osc.Buy(or.UserID, or.CoinID, or.Amount)
	if err != nil {
		return c.JSON(http.StatusBadRequest, NewResponse(math.MinInt, err))
	}
	return c.JSON(http.StatusOK, NewResponse(orderID, nil))
}

func (oh *OrderHandler) Sell(c echo.Context) error {
	or := new(OrderRequest)
	if err := c.Bind(or); err != nil {
		return c.JSON(http.StatusBadRequest, NewResponse(math.MinInt, err))
	}
	if or.Amount <= 0 || or.UserID <= 0 || or.CoinID <= 0 {
		return c.JSON(http.StatusBadRequest, NewResponse(math.MinInt, errors.New("invalid request")))
	}
	orderID, err := oh.osc.Sell(or.UserID, or.CoinID, or.Amount)
	if err != nil {
		return c.JSON(http.StatusBadRequest, NewResponse(math.MinInt, err))
	}
	return c.JSON(http.StatusOK, NewResponse(orderID, nil))
}

func (oh *OrderHandler) Cancel(c echo.Context) error {
	cr := new(CancelRequest)
	if err := c.Bind(cr); err != nil {
		return c.JSON(http.StatusBadRequest, NewResponse(math.MinInt, err))
	}
	if cr.UserID <= 0 || cr.OrderID <= 0 {
		return c.JSON(http.StatusBadRequest, NewResponse(math.MinInt, errors.New("invalid request")))
	}
	err := oh.osc.Cancel(cr.UserID, cr.OrderID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, NewResponse(math.MinInt, err))
	}
	return c.JSON(http.StatusOK, NewResponse(math.MinInt, nil))
}
