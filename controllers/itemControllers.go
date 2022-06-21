package controllers

import (
	"assignment2_/database"
	"assignment2_/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetItems(c *gin.Context) {
	item := []models.Item{}
	db := database.GetDB()
	err := db.Find(&item).Error

	if err != nil {
		fmt.Println("Error getting item data", err.Error())
		return
	}

	c.JSON(http.StatusOK, &item)
}

func AddItem(c *gin.Context) {
	db := database.GetDB()

	newItem := models.Item{}
	var result gin.H

	c.BindJSON(&newItem)
	err := db.Create(&newItem).Error

	if err != nil {
		fmt.Println("Error adding new item", err)
		return
	}

	result = gin.H{
		"status": "Item Added",
		"result": newItem,
	}

	c.JSON(http.StatusCreated, result)
}

func UpdateItem(c *gin.Context) {
	db := database.GetDB()

	newItem := models.Item{}
	var result gin.H

	err := db.Where("id = ?", c.Param("id")).First(&newItem).Error
	c.BindJSON(&newItem)
	db.Save(&newItem)

	if err != nil {
		fmt.Println("Error updating item data", err)
		return
	}

	result = gin.H{
		"status": "Item Updated",
		"result": newItem,
	}

	c.JSON(http.StatusCreated, result)
}

func DeleteItem(c *gin.Context) {
	db := database.GetDB()

	item := models.Item{}

	err := db.Where("id = ?", c.Param("id")).Delete(&item).Error

	if err != nil {
		fmt.Println("Error Delete item data", err)
		return
	}

	result := "Item Deleted"

	c.JSON(http.StatusOK, result)
}