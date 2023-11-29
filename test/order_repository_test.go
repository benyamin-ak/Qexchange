package test

import (
	"Qexchange/internal/core/ordering/models"
	"Qexchange/internal/infrastructure/ordering"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand = rand.NewSource(time.Now().UnixNano())
var r1 = rand.New(seededRand)

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[r1.Intn(len(charset))]
	}
	return string(b)
}

func String(length int) string {
	return StringWithCharset(length, charset)
}

func Int(min int, max int) int {
	return r1.Intn(max-min) + min
}

func Float(min float64, max float64) float64 {
	return r1.Float64()*(max-min) + min
}

var or *ordering.OrderRepository

func DBSetup() *ordering.OrderRepository {
	or := ordering.NewOrderRepository()
	dsn := "root:963233@tcp(127.0.0.1:3306)/Qexchange?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	or.DB = DB
	return or
}

var userList = make([]models.User, 101)
var coinList = make([]struct {
	ID   int    `gorm:"primaryKey;column:id"`
	name string `gorm:"column:name"`
}, 11)
var assetList = make([]struct {
	userID   int     `gorm:"foreignKey:UserRefer;column:user_id"`
	coinID   int     `gorm:"foreignKey:CoinRefer;column:coin_id"`
	quantity float64 `gorm:"column:quantity"`
}, 1001)
var commissionList = make([]struct {
	coinID int     `gorm:"foreignKey:coins.id;column:coin_id"`
	side   string  `gorm:"column:side"`
	rate   float64 `gorm:"column:rate"`
}, 11)

func InitFakeData(or *ordering.OrderRepository) {
	//initfake users
	for i := 1; i < 101; i++ {
		userList[i].FirstName = String(10)
		userList[i].LastName = String(10)
		userList[i].Age = Int(18, 80)
		userList[i].NationalCode = String(10)
		userList[i].AuthenticationLevel = String(1)
		userList[i].Phone = String(10)
		userList[i].Email = String(10)
		userList[i].AccountType = String(1)
		userList[i].Username = String(10)
		userList[i].Password = String(10)
		u := userList[i]
		or.DB.Exec("INSERT INTO users (id,first_name, last_name,age,national_code,authentication_level,phone,email,account_type,username,password) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);",
			i, u.FirstName, u.LastName, u.Age, u.NationalCode, u.AuthenticationLevel, u.Phone, u.Email, u.AccountType, u.Username, u.Password)
	}
	//initfake coins

	for i := 1; i < 11; i++ {
		coinList[i].ID = i
		coinList[i].name = String(10)
		or.DB.Exec("INSERT INTO coins (id, name) VALUES (?, ?);", coinList[i].ID, coinList[i].name)
	}
	//initfake assets

	for i := 0; i < 100; i++ {
		for j := 0; j < 10; j++ {
			assetList[i*10+j].userID = i + 1
			assetList[i*10+j].coinID = j + 1
			assetList[i*10+j].quantity = Float(0, 1000)
			or.DB.Exec("INSERT INTO assets (user_id, coin_id, quantity) VALUES (?, ?, ?);", i+1, j+1, assetList[i*10+j].quantity)
		}
	}
	//initfake commision
	for i := 1; i < 11; i++ {
		commissionList[i].coinID = coinList[i].ID
		commissionList[i].rate = Float(0, 1)
		or.DB.Exec("INSERT INTO commissions (coin_id ,rate) VALUES (?, ?);", commissionList[i].coinID, commissionList[i].rate)
	}

	//initfake price
	for i := 1; i < 11; i++ {
		or.DB.Exec("INSERT INTO mock_coin_prices (coin_id ,price) VALUES (?, ?);", i, Float(1, 1000))
	}
}

func Test(t *testing.T) {
	or = DBSetup()
	InitFakeData(or)
	t.Run("TestGetUserBalance", TestGetUserBalance)
	t.Run("TestGetUserBalanceWithInvalidUserID", TestGetUserBalanceWithInvalidUserID)
	t.Run("TestGetUserBalanceWithInvalidCoinID", TestGetUserBalanceWithInvalidCoinID)
	t.Run("TestGetCoinPrice", TestGetCoinPrice)
	t.Run("TestGetCoinPriceWithInvalidCoinID", TestGetCoinPriceWithInvalidCoinID)
	t.Run("TestGetCoinCommission", TestGetCoinCommission)
	t.Run("TestGetCoinCommissionWithInvalidCoinID", TestGetCoinCommissionWithInvalidCoinID)
	t.Run("TestCreateOrder", TestCreateOrder)
	t.Run("TestCreateOrderWithInvalidUserID", TestCreateOrderWithInvalidUserID)
	t.Run("TestCreateOrderWithInvalidCoinID", TestCreateOrderWithInvalidCoinID)
	t.Run("TestSubmitOrder", TestSubmitOrder)
	t.Run("TestChangeOrderStatus", TestChangeOrderStatus)
	t.Run("TestValidateOrderBelongToUser", TestValidateOrderBelongToUser)
}

func TestGetUserBalance(t *testing.T) {
	for i := 0; i < 100; i++ {
		for j := 0; j < 10; j++ {
			asset, err := or.GetUserBalance(i+1, j+1)
			assert.Nil(t, err)
			assert.Equal(t, assetList[i*10+j].quantity, asset)
		}
	}
}

func TestGetUserBalanceWithInvalidUserID(t *testing.T) {
	_, err := or.GetUserBalance(101, 1)
	assert.NotNil(t, err)
}

func TestGetUserBalanceWithInvalidCoinID(t *testing.T) {
	_, err := or.GetUserBalance(1, 11)
	assert.NotNil(t, err)
}

func TestGetCoinPrice(t *testing.T) {
	for i := 1; i < 11; i++ {
		price, err := or.GetCoinPrice(i)
		assert.Nil(t, err)
		assert.Equal(t, 0.0, price)
	}
}

func TestGetCoinPriceWithInvalidCoinID(t *testing.T) {
	_, err := or.GetCoinPrice(11)
	assert.NotNil(t, err)
}

func TestGetCoinCommission(t *testing.T) {
	for i := 1; i < 11; i++ {
		commission, err := or.GetCoinCommission(i)
		assert.Nil(t, err)
		assert.Equal(t, commissionList[i].rate, commission)
	}
}

func TestGetCoinCommissionWithInvalidCoinID(t *testing.T) {
	_, err := or.GetCoinCommission(11)
	assert.NotNil(t, err)
}

func TestCreateOrder(t *testing.T) {
	for i := 1; i < 101; i++ {
		o := models.Order{
			OrderID:  i,
			UserID:   i,
			Side:     "buy",
			Quantity: Float(1, 1000),
			CoinID:   Int(1, 10),
			Status:   "active",
		}
		price, _ := or.GetCoinPrice(o.CoinID)
		o.Price = price
		_, err := or.CreateOrder(&o)
		assert.Nil(t, err)
		qo := &models.Order{}
		or.DB.Table("orders").Where("id = ?", i).Scan(qo)
		assert.Equal(t, o, *qo)
	}
}

func TestCreateOrderWithInvalidUserID(t *testing.T) {
	o := models.Order{
		OrderID:  101,
		UserID:   101,
		Side:     "buy",
		Quantity: Float(1, 1000),
		CoinID:   Int(1, 10),
		Status:   "active",
	}
	price, _ := or.GetCoinPrice(o.CoinID)
	o.Price = price
	_, err := or.CreateOrder(&o)
	assert.NotNil(t, err)
}

func TestCreateOrderWithInvalidCoinID(t *testing.T) {
	o := models.Order{
		OrderID:  101,
		UserID:   1,
		Side:     "buy",
		Quantity: Float(1, 1000),
		CoinID:   11,
		Status:   "active",
	}
	price, _ := or.GetCoinPrice(o.CoinID)
	o.Price = price
	_, err := or.CreateOrder(&o)
	assert.NotNil(t, err)
}

func TestSubmitOrder(t *testing.T) {
	for i := 1; i < 101; i++ {
		o := models.Order{
			OrderID:  i,
			UserID:   i,
			Side:     "buy",
			Quantity: Float(1, 1000),
			CoinID:   Int(1, 10),
			Status:   "active",
		}
		price, _ := or.GetCoinPrice(o.CoinID)
		o.Price = price
		or.SubmitOrder(&o)
		qo := &models.Order{}
		or.DB.Table("orders").Where("id = ?", i).Scan(qo)
		assert.Equal(t, "completed", qo.Status)
	}
}

func TestChangeOrderStatus(t *testing.T) {
	for i := 1; i < 101; i++ {
		o := models.Order{
			OrderID: i,
			UserID:  i,
		}
		str := []string{"active", "completed", "cancelled"}
		randstr := str[Int(0, 3)]
		or.ChangeOrderStatus(&o, randstr)
		qo := &models.Order{}
		or.DB.Table("orders").Where("id = ?", i).Scan(qo)
		assert.Equal(t, randstr, qo.Status)
	}
}

func TestValidateOrderBelongToUser(t *testing.T) {
	for i := 1; i < 101; i++ {
		o := models.Order{
			OrderID: i,
		}
		err := or.ValidateOrderBelongToUser(&o, i)
		assert.Nil(t, err)
	}
}
