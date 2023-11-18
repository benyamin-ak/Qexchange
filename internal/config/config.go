package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Database Database
	// PaymentGateway PaymentGateway
}

type Database struct {
	Driver   string
	Host     string
	Port     int
	DB       string
	Username string
	Password string
	Charset  string
}

// type PaymentGateway struct {
// 	TerminalId       string
// 	MerchantId    string
// }

func Init(param Params) (*Config, error) {
	viper.SetConfigType(param.FileType)
	viper.AddConfigPath(param.FilePath)
	viper.SetConfigFile("./config/config.yaml")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %s", err)
	}

	database := &Database{
		Driver:   viper.GetString("database.driver"),
		Host:     viper.GetString("database.host"),
		Port:     viper.GetInt("database.port"),
		Username: viper.GetString("database.username"),
		Password: viper.GetString("database.password"),
		Charset:  viper.GetString("database.chaset"),
		DB:       viper.GetString("database.db"),
	}

	// paymentGateway := &PaymentGateway{
	// 	TerminalId:       viper.GetString("payment_gateway.TerminalId"),
	// 	APIKey:    viper.GetString("payment_gateway.MerchantId"),
	// }

	return &Config{
		Database: *database,
		// PaymentGateway: *paymentGateway,
	}, nil
}
