package main

import (
	"cursogo/templates"
	"cursogo/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Static("/static", "static")
	e.GET("/", func(c echo.Context) error {
		return utils.Render(c, http.StatusOK, templates.Hello("World"))
	})
	e.Logger.Fatal(e.Start(":3000"))
}
