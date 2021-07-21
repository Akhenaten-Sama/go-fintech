package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)
type User struct {
	gorm.Model
	Username string
	Email string
	Password string
}

type Account struct {
	gorm.Model
	Type string
	Name string
	Balance uint
	UserID uint
}


type ResponseAccount struct {
		ID uint
		Name string 
		Balance int

}

type ResponseUser struct {
	ID uint
	Username string
	Email string
	Accounts []ResponseAccount
}