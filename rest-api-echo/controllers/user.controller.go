package controllers

import (
	"cursogo/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, models.Users)
}

func GetUser(c echo.Context) error {
	userId := c.Param("id")

	for _, user := range models.Users {
		if user.Id == userId {
			return c.JSON(http.StatusOK, user)
		}
	}

	return c.String(http.StatusNotFound, "User not found")
}
func CreateUser(c echo.Context) error {
	var newUser models.User

	err := c.Bind(&newUser)

	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid data")
	}

	models.Users = append(models.Users, newUser)

	return c.JSON(http.StatusCreated, newUser)
}

func UpdateUser(c echo.Context) error {
	userId := c.Param("id")

	var newUser models.User
	err := c.Bind(&newUser)

	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid data")
	}

	for i, user := range models.Users {
		if user.Id == userId {
			models.Users[i].IsActive = newUser.IsActive
			models.Users[i].Name = newUser.Name

			return c.JSON(http.StatusOK, models.Users[i])
		}
	}

	return c.String(http.StatusNotFound, "User not found")
}

func PartialUpdateUser(c echo.Context) error {
	return c.String(http.StatusOK, "Hello Wold")
}

func DeleteUser(c echo.Context) error {
	var userId string = c.Param("id")

	for i, user := range models.Users {
		if user.Id == userId {
			models.Users = append(models.Users[:i], models.Users[i+1:]...)
			return c.NoContent(http.StatusNoContent)
		}
	}

	return c.String(http.StatusNotFound, "User not found")
}
