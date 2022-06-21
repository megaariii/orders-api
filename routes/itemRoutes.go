package routes

import (
	"assignment2_/controllers"

	"github.com/gin-gonic/gin"
)
func StartServer() *gin.Engine {
	router := gin.Default()

	router.GET("/items", controllers.GetItems)
	router.POST("/item", controllers.CreateItem)
	router.PUT("/item/:id", controllers.UpdateItem)
	router.DELETE("/item/:id", controllers.DeleteItem)

	return router
}