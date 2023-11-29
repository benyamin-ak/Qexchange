package models

type User struct {
	ID                  int    `gorm:"primaryKey;column:id"`
	FirstName           string `gorm:"column:first_name"`
	LastName            string `gorm:"column:last_name"`
	Age                 int    `gorm:"column:age"`
	NationalCode        string `gorm:"column:national_code"`
	AuthenticationLevel string `gorm:"column:authentication_level"`
	Phone               string `gorm:"column:phone"`
	Email               string `gorm:"column:email"`
	AccountType         string `gorm:"column:account_type"`
	Username            string `gorm:"column:username"`
	Password            string `gorm:"column:password"`
}
