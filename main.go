package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func helloWorldHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World")
}

func pingHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Pong!")
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", helloWorldHandler)

	e.GET("/ping", pingHandler)

	e.Logger.Fatal(e.Start((":3000")))
}
