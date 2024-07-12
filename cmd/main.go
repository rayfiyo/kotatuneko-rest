package main

import (
	"github.com/labstack/echo/v4"
	_ "github.com/rayfiyo/kotatuneko-rest/docs"
	"github.com/rayfiyo/kotatuneko-rest/internal/app/handler/hello"
	echoswagger "github.com/swaggo/echo-swagger"
)

// @title			Kotatuneko REST API
// @version			1.0
// @description		こたつねこの REST API
// @termsOfService	http://localhost:8080/

// @license.name	BSD-3-Clause license
// @license.url		https://github.com/rayfiyo/kotatuneko-rest?tab=BSD-3-Clause-1-ov-file

// @host						localhost:8080
// @BasePath					/
// @SecurityDefinitions.apiKey	Bearer
// @in							header
// @name						Authorization

func main() {
	e := echo.New()

	e.GET("/hello", hello.Hello)
	e.GET("/swagger/*", echoswagger.WrapHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
