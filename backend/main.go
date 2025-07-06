package main

import (
	"fmt"

	"github.com/endr0id/locksmith/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	handler.RegisterRoutes(e)

	port := 8080
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}