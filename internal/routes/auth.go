package routes

import (
	"learn-go/internal/handler"

	"github.com/labstack/echo/v4"
)

func AuthRoutes(e *echo.Echo, h *handler.AuthHandler) {
	e.POST("/signup", h.Signup)
	e.POST("/login", h.Login)
}
