package main

import (
	"learn-go/internal/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	routes.HealthRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))
}
