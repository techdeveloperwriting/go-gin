package db

import (
	"fmt"
	"go_gin_crud/models"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var Db *gorm.DB

func init() {
	fmt.Println("Init ")
	err := godotenv.Load("local.env")
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
	VAL := os.Getenv("VAL")
	PASS := os.Getenv("PASS")
	HOST := os.Getenv("HOST")
	PORT := os.Getenv("PORT")
	DBNAME := os.Getenv("DBNAME")
	fmt.Println("Pas", VAL, PASS)
	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", VAL, PASS, HOST, PORT, DBNAME)
	Db, err = gorm.Open("mysql", URL)
	if err != nil {
		panic(err.Error())
	}
	Db.AutoMigrate(&models.User{})
	Db.AutoMigrate(&models.Login{})
	fmt.Println("Init ", Db)
}
func SetupDB() *gorm.DB {
	return Db
}
