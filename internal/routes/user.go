package routes

import (
	"learn-go/internal/config"
	"learn-go/internal/handler"
	"learn-go/internal/middleware"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Echo, h *handler.UserHandler, cfg *config.Config) {
	e.GET("/me", h.Me, middleware.JWTCookieMiddleware(cfg))
}
