package routes

import (
	"assignment2_/controllers"

	"github.com/gin-gonic/gin"
)
func StartServer() *gin.Engine {
	router := gin.Default()

	router.GET("/orders", controllers.GetOrders)
	router.GET("/order/:id", controllers.GetOrderById)
	router.POST("/order", controllers.CreateOrder)
	router.PUT("/order/:id", controllers.UpdateOrder)
	router.DELETE("/order/:id", controllers.DeleteOrder)

	router.GET("/items", controllers.GetItems)
	router.POST("/item", controllers.AddItem)
	router.PUT("/item/:id", controllers.UpdateItem)
	router.DELETE("/item/:id", controllers.DeleteItem)

	return router
}