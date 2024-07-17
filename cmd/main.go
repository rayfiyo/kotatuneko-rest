package main

import (
	"log"

	"github.com/labstack/echo/v4/middleware"
	_ "github.com/rayfiyo/kotatuneko-rest/docs"
	"github.com/rayfiyo/kotatuneko-rest/internal/domain/service/user"
	"github.com/rayfiyo/kotatuneko-rest/internal/infrastructure/db"
	"github.com/rayfiyo/kotatuneko-rest/internal/infrastructure/repository"
	"github.com/rayfiyo/kotatuneko-rest/internal/interface/handler"
	"github.com/rayfiyo/kotatuneko-rest/internal/interface/router"
	"github.com/rayfiyo/kotatuneko-rest/internal/usecase"

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
	database := db.New()

	userRepo := repository.NewUserRepository(database)
	userService := service.NewUserService(userRepo)
	userUsecase := usecase.NewUserUsecase(userRepo, userService)
	userHandler := handler.NewUserHandler(userUsecase)

	e := router.NewRouter(userHandler)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/swagger/*", echoswagger.WrapHandler)

	log.Println("Server is running at http://localhost:8080")
	e.Logger.Fatal(e.Start(":8080"))
}
