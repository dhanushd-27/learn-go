package main

import (
	"learn-go/internal/service/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/health", func(c echo.Context) error {
		h := &response.Health{
			Message: "OK",
		}
		return c.JSON(http.StatusOK, h)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
