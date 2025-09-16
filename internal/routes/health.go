package routes

import (
	"learn-go/internal/handler"

	"github.com/labstack/echo/v4"
)

func HealthRoutes(e *echo.Echo) {
	e.GET("/health", handler.Health)
}
