package main

import (
	"net/http"
	"rest-api/database"
	"rest-api/models"
	"rest-api/router"
)

func main() {
	database.ConnectDB()

	database.DB.AutoMigrate(models.User{})
	database.DB.AutoMigrate(models.Task{})
	Router := router.Init()

	http.ListenAndServe(":3000", Router)
}
