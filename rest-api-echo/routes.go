package main

import (
	"cursogo/controllers"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	users := e.Group("/users")

	users.GET("/", controllers.GetUsers)
	users.GET("/:id", controllers.GetUser)
	users.POST("/", controllers.CreateUser)
	users.PUT("/:id", controllers.UpdateUser)
	users.PATCH("/:id", controllers.PartialUpdateUser)
	users.DELETE("/:id", controllers.DeleteUser)
}
