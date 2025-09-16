package main

import (
	"learn-go/internal/container"
	"learn-go/internal/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	c, err := container.NewContainer()
	if err != nil {
		e.Logger.Fatal(err)
	}

	routes.HealthRoutes(e)

	e.Logger.Fatal(e.Start(c.Config.PORT))
}
