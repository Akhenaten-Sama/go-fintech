package users

import (
	"time"
	
	"github.com/Akhenaten-Sama/go-fintech/helpers"
	"github.com/Akhenaten-Sama/go-fintech/models"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func Login(username string, password string) map[string]interface{} {
 db:= helpers.ConnectDB()

 user := &models.User{}
 if db.Where("username = ?", username).First(&user).RecordNotFound(){
	 return map[string]interface{}{"message": "User not found"}
 }

 passErr:= bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

 if passErr == bcrypt.ErrMismatchedHashAndPassword && passErr !=nil {

	return map[string]interface{}{"message": "Wrong Password"}
 }


 accounts:=[]models.ResponseAccount{}
 db.Table("account").Select("id, name, balance").Where("user_id = ?", user.ID).scan(&accounts)

 responseUser:= &models.ResponseUser{
	 ID: user.ID,
	 Username:  user.Username,
	 Email: user.Email,
	 Accounts: accounts,
 }

 defer db.Close()

	tokenContent := jwt.MapClaims{
		"user_id": user.ID,
		"expiry" : time.Now().Add(time.Minute *60).Unix(),
	}

 
}