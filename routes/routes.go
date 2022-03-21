package routes

import (
	"go_gin_crud/handlers"

	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	router := gin.Default()
	router.POST("/create", handlers.CreateUser)
	//  router.POST("/create", handlers.CreateUser)
	router.PUT("/update", handlers.UpdateUser)
	router.GET("/:id", handlers.GetUser)
	router.POST("/delete/:id", handlers.DeleteUser)
	router.POST("/login", handlers.LoginUser)
	return router
}
