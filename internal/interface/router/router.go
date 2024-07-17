package router

import (
	"github.com/labstack/echo/v4"
	"github.com/rayfiyo/kotatuneko-rest/internal/interface/handler"
)

func NewRouter(userHandler *handler.UserHandler) *echo.Echo {
	e := echo.New()

	e.POST("/regiseter", userHandler.RegisterUser)
	e.POST("/login", userHandler.LoginUser)
	e.PUT("/update", userHandler.UpdateUser)
	e.GET("/rankign", userHandler.GetRanking)

	return e
}
