package main

import (
	"context"
	"learn-go/internal/config"
	"learn-go/internal/container"
	"learn-go/internal/handler"
	"learn-go/internal/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	cfg, err := config.Load()
	if err != nil {
		e.Logger.Fatal(err)
	}

	ctx := context.Background()
	c, err := container.NewContainer(ctx, cfg)
	if err != nil {
		e.Logger.Fatal(err)
	}
	defer c.Close()

	// Handlers
	authHandler := handler.NewAuthHandler(c.GetDB(), c.GetConfig(), c.GetQuery())

	// Routes
	routes.HealthRoutes(e)
	routes.AuthRoutes(e, authHandler)

	e.Logger.Fatal(e.Start(":" + c.GetConfig().PORT))
}
