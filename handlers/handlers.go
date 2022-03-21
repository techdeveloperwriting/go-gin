package handlers

import (
	"fmt"
	"go_gin_crud/db"
	"go_gin_crud/models"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// type User struct {
// 	Name string
// }
var Db *gorm.DB

func Handlers() {

}

func init() {
	fmt.Println("Init Handler")
	Db = db.SetupDB()
	fmt.Println("Init Handler", Db)
}
func CreateUser(c *gin.Context) {
	fmt.Println("Init ", Db)
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Create User
	Db.Create((&input))
	c.JSON(http.StatusOK, gin.H{"data": input})
}
func UpdateUser(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Update User

	Db.Model(&input).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": input})
}

func GetUser(c *gin.Context) {
	fmt.Println("Get user : ")
	var input models.User

	if err := Db.Where("Id = ?", c.Param("id")).First(&input).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": input})
}

func DeleteUser(c *gin.Context) {
	var input models.User

	if err := Db.Where("Id = ?", c.Param("id")).Delete(&input).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "Record"})
}

var user = models.Login{
	ID:       1,
	Username: "username",
	Password: "password",
}

func LoginUser(c *gin.Context) {

	var input models.Login

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	fmt.Println("LoginUser : ", input.Username)
	fmt.Println("LoginUser : ", user.Username)
	fmt.Println("LoginUser : ", user.Password)
	fmt.Println("Password : ", input.Password)
	//compare the user from the request, with the one we defined:
	if user.Username != input.Username || user.Password != input.Password {
		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		return
	}
	token, err := CreateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK, token)
}

func CreateToken(userId uint64) (string, error) {
	var err error
	//Creating Access Token
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") //this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userId
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}
