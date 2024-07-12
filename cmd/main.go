package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rayfiyo/kotatuneko-rest/internal/app/handler/hello"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", hello.Hello)

	e.Logger.Fatal(e.Start(":8080"))
}
