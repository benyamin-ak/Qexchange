package main

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID                  uint `gorm:"primaryKey"`
	FirstName           string
	LastName            string
	Age                 int
	NationalCode        string
	AuthenticationLevel string
	Phone               string
	Email               string
	AccountType         string
	Username            string
	Password            string
}

type BankAccount struct {
	ID            uint `gorm:"primaryKey"`
	UserID        uint
	User          User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CardNumber    string
	AccountNumber string
}

type AuthenticationLevel struct {
	ID    uint `gorm:"primaryKey"`
	Level string
}

type UserStatus struct {
	ID            uint `gorm:"primaryKey"`
	UserID        uint
	User          User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Endtime       time.Time `gorm:"type:timestamp"`
	AccountStatus string
}

type Order struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	User      User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Side      string
	Quantity  float64
	Price     float64
	CoinID    uint
	Coin      Coin      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	timestamp time.Time `gorm:"type:timestamp"`
	Status    string
}

type Admin struct {
	ID        uint `gorm:"primaryKey"`
	FirstName string
	LastName  string
	Email     string
	Username  string
	Password  string
}

type Ticket struct {
	ID      uint `gorm:"primaryKey"`
	UserID  uint
	User    User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Detail  string
	Status  string
	AdminID uint
	Admin   Admin `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
type Coin struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}

type MockCoinPrice struct {
	CoinID uint
	Coin   Coin `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Price  float64
}
type Asset struct {
	UserID   uint
	User     User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CoinID   uint
	Coin     Coin `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Quantity float64
}

type Commission struct {
	CoinID uint
	Coin   Coin `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	side   string
	rate   float64
}

type Withdraw_deposit_money struct {
	ID            uint `gorm:"primaryKey"`
	UserID        uint
	User          User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Type          string
	Amount        float64
	BankAccountID uint
	BankAccount   BankAccount `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func main() {
	dsn := "root:Ab@12344321@tcp(localhost:3306)/qexchange?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}
	// Save the 'db' variable for later use.

	// Set a maximum time for the database ping to complete
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	sqlDB.SetConnMaxLifetime(time.Minute * 3)

	// Try to ping the database
	err = sqlDB.Ping()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	// AutoMigrate all models
	err = db.AutoMigrate(&User{}, &BankAccount{}, &AuthenticationLevel{}, &UserStatus{}, &Order{}, &Admin{}, &Coin{}, &Mock_Coin_Price{}, &Asset{}, &Commission{}, &Ticket{}, &Withdraw_deposit_money{})
	if err != nil {
		panic(err)
	}

	fmt.Println("Database migration successful!")
	defer sqlDB.Close()
}
