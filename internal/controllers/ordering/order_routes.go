package ordering

import echo "github.com/labstack/echo/v4"

func DefineRoutes(e *echo.Echo) {
	oh := NewOrderHandler()
	oah := NewAutoOrderHandler()
	e.POST("/buy", oh.Buy)
	e.POST("/sell", oh.Sell)
	e.PUT("/cancel", oh.Cancel)
	e.POST("/auto_order", oah.CreateNewAutoOrderHandler)
}
