package database

import (
	"assignment2_/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host		= "localhost"
	user		= "developer"
	password	= "supersecretpassword"
	dbPort		= "5432"
	dbname		= "orderdb"
	db 			*gorm.DB
	err			error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, dbPort)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database :", err)
	}

	db.Debug().AutoMigrate(models.Item{}, models.Order{})
}

func GetDB() *gorm.DB {
	return db
}