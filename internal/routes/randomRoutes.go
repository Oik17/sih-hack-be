package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/oik17/sih-agrihealth/internal/controllers"
)

func RandomRoutes(e *echo.Echo) {

	e.POST("/upload", controllers.UploadFilesToS3)
	e.POST("/translate", controllers.Translate)
	e.GET("/news", controllers.NewsControllers)
}
