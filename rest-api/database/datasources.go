package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	DB, err = gorm.Open(postgres.Open("host=localhost user=postgres password=test dbname=restapi port=5432 sslmode=disable"))

	if err != nil {
		log.Fatal("Error connecting to database " + err.Error())
		return
	}
	log.Println("Database connected")
}
