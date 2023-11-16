package controllers

import (
	"Qexchange/internal/core/contracts"
	"Qexchange/internal/core/services"
	"net/http"
	"strconv"

	echo "github.com/labstack/echo/v4"
)

type OrderRequest struct {
	UserID int     `json:"user_id"`
	CoinID int     `json:"coin_id"`
	Amount float64 `json:"amount"`
}

type OrderHandler struct {
	osc contracts.OrderCoreContract
}

func NewOrderHandler() *OrderHandler {
	return &OrderHandler{
		osc: services.NewOrderService(),
	}
}

func DefineRoutes(e *echo.Echo) {
	oh := NewOrderHandler()
	e.POST("/buy", oh.Buy)
	e.POST("/sell", oh.Sell)
	e.POST("/cancel", oh.Cancel)
}

func (oh *OrderHandler) Buy(c echo.Context) error {
	or := new(OrderRequest)
	if err := c.Bind(or); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	orderID, err := oh.osc.Buy(or.CoinID, or.Amount)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, "orderID: "+strconv.Itoa(orderID))
}

func (oh *OrderHandler) Sell(c echo.Context) error {
	or := new(OrderRequest)
	if err := c.Bind(or); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	orderID, err := oh.osc.Sell(or.CoinID, or.Amount)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, "orderID: "+strconv.Itoa(orderID))
}

func (oh *OrderHandler) Cancel(c echo.Context) error {
	return nil
}
