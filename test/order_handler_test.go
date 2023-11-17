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
	requests := []string{`{"user_id": 1,"coin_id": 1}`, `{"user_id": 1,"amount": 1}`, `{"coin_id": 1,"amount": 1}`}
	responses := make([]*controllers.Response, len(requests))
	for i, req := range requests {
		resp, err := invalidReqGenertor(req)
		assert.Nil(t, err)
		responses[i] = &controllers.Response{}
		err = json.NewDecoder(resp.Body).Decode(responses[i])
		assert.Nil(t, err)
		assert.Equal(t, math.MinInt, responses[i].OrderID)
		assert.Equal(t, "invalid request", responses[i].Error)
	}
}
func invalidReqGenertor(req string) (*http.Response, error) {
	return http.DefaultClient.Post("http://localhost:8080/buy", "application/json", strings.NewReader(req))
}

func TestOrderHandlerSellInvalidRequest(t *testing.T) {
	requests := []string{`{"user_id": 1,"coin_id": 1}`, `{"user_id": 1,"amount": 1}`, `{"coin_id": 1,"amount": 1}`}
	responses := make([]*controllers.Response, len(requests))
	for i, req := range requests {
		resp, err := invalidReqGenertor(req)
		assert.Nil(t, err)
		responses[i] = &controllers.Response{}
		err = json.NewDecoder(resp.Body).Decode(responses[i])
		assert.Nil(t, err)
		assert.Equal(t, math.MinInt, responses[i].OrderID)
		assert.Equal(t, "invalid request", responses[i].Error)
	}
}

func TestOrderHandlerCancelInvalidRequest(t *testing.T) {
	requests := []string{`{"user_id": 1,"order_id": 1}`, `{"user_id": 1,"user_password": "123456"}`, `{"order_id": 1,"user_password": "123456"}`}
	responses := make([]*controllers.Response, len(requests))
	for i, req := range requests {
		resp, err := invalidReqGenertor(req)
		assert.Nil(t, err)
		responses[i] = &controllers.Response{}
		err = json.NewDecoder(resp.Body).Decode(responses[i])
		assert.Nil(t, err)
		assert.Equal(t, math.MinInt, responses[i].OrderID)
		assert.Equal(t, "invalid request", responses[i].Error)
	}
}
