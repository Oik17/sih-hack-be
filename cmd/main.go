package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/oik17/sih-agrihealth/internal/database"
	"github.com/oik17/sih-agrihealth/internal/routes"
)

func main() {

	database.Connect()
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "pong")
	})
	routes.RandomRoutes(e)
	e.Start(":8080")
}
