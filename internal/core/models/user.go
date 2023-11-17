package models

type User struct {
	ID                    int
	firstName             string
	lastName              string
	age                   int
	nationalCode          string
	authenticationLevelID string
	mobileNumber          string
	email                 string
	accountTypeID         string
	username              string
	password              string
}
