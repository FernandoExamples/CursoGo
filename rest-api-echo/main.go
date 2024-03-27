package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func handleIndex(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "Hello, World!"})
}

func main() {
	e := echo.New()

	e.GET("/", handleIndex)

	e.Logger.Fatal(e.Start(":3000"))
}
