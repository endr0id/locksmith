package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "locksmith is alive")
	})

	port := 8080
	fmt.Printf("Starting server on port %d...\n", port)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}