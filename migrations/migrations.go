package migrations

import (

	"github.com/Akhenaten-Sama/go-fintech/helpers"
	"github.com/Akhenaten-Sama/go-fintech/models"
)





func CreateAccounts(){
	db:= helpers.ConnectDB()
	users :=[2]models.User{
		{Username:"Efukunle", Email:"olalekan@gmail.com"},
		{Username:"Olanrewaju", Email:"olanike@gmail.com"},
	}

	for i := 0; i <len(users); i++ {
		generatePassword := helpers.HashAndSalt([]byte(users[i].Username))
		user := models.User{Username:users[i].Username, Email:users[i].Email, Password: generatePassword}
         db.Create(&user)

		 account := models.Account{Type:"Savings", Name:users[i].Username, Balance: uint(10000 * int(i+1)), UserID: user.ID}
		 db.Create(&account)
	 
	}
	defer db.Close()
}

func Migrate(){
	db:= helpers.ConnectDB()
	db.AutoMigrate(&models.User{}, &models.Account{})
	defer db.Close()

	CreateAccounts()
}