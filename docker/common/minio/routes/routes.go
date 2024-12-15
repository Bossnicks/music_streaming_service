package routes

import (
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, fileHandler *handler.FileHandler) {
	e.POST("/upload", fileHandler.UploadFile)
	e.GET("/files/:id", fileHandler.GetFile)
	e.DELETE("/files/:id", fileHandler.DeleteFile)
}
