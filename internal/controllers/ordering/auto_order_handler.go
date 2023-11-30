package ordering

import (
	"Qexchange/internal/core/ordering"
	"errors"
	"math"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AutoOrderHandler struct {
	oac ordering.AutoOrderContract
}

func NewAutoOrderHandler() *AutoOrderHandler {
	return &AutoOrderHandler{
		oac: ordering.NewAutoOrderService(),
	}
}

func (aoh *AutoOrderHandler) CreateNewAutoOrderHandler(c echo.Context) error {
	aor := new(AutoOrderRequest)
	if err := c.Bind(aor); err != nil {
		return c.JSON(http.StatusBadRequest, NewOrderResponse(math.MinInt, err))
	}
	if aor.Quantity <= 0 || aor.UserID <= 0 || aor.CoinID <= 0 || aor.PTS <= 0 || (aor.Side != "buy" && aor.Side != "sell") {
		return c.JSON(http.StatusBadRequest, NewOrderResponse(math.MinInt, errors.New("invalid request")))
	}
	aoh.oac.StartPolling(aor.UserID, aor.CoinID, aor.Quantity, aor.Side, aor.PTS)
	return c.JSON(http.StatusOK, NewOrderResponse(0, nil))
}
