package models

type User struct {
	ID                    int `gorm:"primaryKey"`
	firstName             string
	lastName              string
	age                   int
	nationalCode          string
	authenticationLevelID string `gorm:"foreignKey:authenticationLevels"`
	mobileNumber          string
	email                 string
	accountTypeID         string
	username              string
	password              string
}
