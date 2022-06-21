package controllers

import (
	"assignment2_/database"
	"assignment2_/models"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetOrders(c *gin.Context) {
	order := []models.Order{}
	db := database.GetDB()
	err := db.Find(&order).Error

	if err != nil {
		fmt.Println("Error getting order data", err.Error())
		return
	}

	c.JSON(http.StatusOK, &order)
}

func GetOrderById(c *gin.Context) {
	order := models.Order{}
	db := database.GetDB()
	
	err := db.Where("id = ?", c.Param("id")).First(&order).Error
	db.Preload("Items").Find(&order)

	if err != nil {
		fmt.Println("Error getting order data", err.Error())
		return
	}

	c.JSON(http.StatusOK, &order)
}

func CreateOrder(c *gin.Context) {
	db := database.GetDB()

	newOrder := models.Order{
		OrderedAt: time.Now(),
	}
	var result gin.H

	c.BindJSON(&newOrder)
	err := db.Create(&newOrder).Error

	if err != nil {
		fmt.Println("Error creating order", err)
		return
	}

	result = gin.H{
		"result": newOrder,
	}

	c.JSON(http.StatusCreated, result)
}

func UpdateOrder(c *gin.Context) {
	db := database.GetDB()

	newOrder := models.Order{
		OrderedAt: time.Now(),
	}
	var result gin.H

	err := db.Where("id = ?", c.Param("id")).First(&newOrder).Error
	c.BindJSON(&newOrder)
	db.Save(&newOrder)

	if err != nil {
		fmt.Println("Error updating order data", err)
		return
	}

	result = gin.H{
		"result": newOrder,
	}

	c.JSON(http.StatusCreated, result)
}

func DeleteOrder(c *gin.Context) {
	db := database.GetDB()

	order := models.Order{}

	err := db.Where("id = ?", c.Param("id")).Delete(&order).Error

	if err != nil {
		fmt.Println("Error Delete order data", err)
		return
	}

	result := "Order Deleted"

	c.JSON(http.StatusOK, result)
}