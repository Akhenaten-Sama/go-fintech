package helpers

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)
 


func HandleErr (err error){
 if err != nil{
	 panic(err.Error())
 }
}



func ConnectDB() *gorm.DB {
	err := godotenv.Load(".env")
	HandleErr(err)
	username := os.Getenv("APP_DB_USERNAME")
pass := os.Getenv("APP_DB_PASSWORD")
host := os.Getenv("APP_DB_HOST")
dbName := os.Getenv("APP_DB_NAME")
port := os.Getenv("APP_DB_PORT")
connects := fmt.Sprintf("host=%s port=%s user=%s "+
"password=%s dbname=%s sslmode=disable", host, port, username, pass, dbName)
	db, err := gorm.Open("postgres", connects)
	HandleErr(err)
	return db
}

func HashAndSalt (password []byte) string{
	hashed, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	HandleErr(err)
	return string(hashed)
}