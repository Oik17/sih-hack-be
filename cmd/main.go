package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oik17/sih-agrihealth/internal/controllers"
	"github.com/oik17/sih-agrihealth/internal/database"
)

func main() {

	database.Connect()
	e := echo.New()
	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "pong")
	})
	e.POST("/upload", controllers.UploadFilesToS3)
	e.GET("/news", controllers.NewsControllers)
	e.Start(":8080")
}
