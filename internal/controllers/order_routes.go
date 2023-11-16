package controllers

import echo "github.com/labstack/echo/v4"

func DefineRoutes(e *echo.Echo) {
	oh := NewOrderHandler()
	e.POST("/buy", oh.Buy)
	e.POST("/sell", oh.Sell)
	e.PUT("/cancel", oh.Cancel)
}
