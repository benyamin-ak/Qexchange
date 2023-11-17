package test

import (
	"Qexchange/internal/controllers"
	"encoding/json"
	"math"
	"net/http"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func Setup() {
	e := echo.New()
	controllers.DefineRoutes(e)
	e.Start(":8080")
}

func TestMain(m *testing.M) {
	Setup()
	m.Run()
}
func TestOrderHandlerBuyInvalidRequest(t *testing.T) {
	oreq := `{"user_id": 1,"coin_id": 1}`
	resp, err := http.DefaultClient.Post("http://localhost:8080/buy", "application/json", strings.NewReader(oreq))
	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	ores := &controllers.Response{}
	err = json.NewDecoder(resp.Body).Decode(ores)
	assert.Nil(t, err)
	assert.Equal(t, math.MinInt, ores.OrderID)
	assert.Equal(t, "invalid request", ores.Error)
}
