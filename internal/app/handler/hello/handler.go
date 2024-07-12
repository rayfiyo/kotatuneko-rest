package hello

import (
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/rayfiyo/kotatuneko-rest/docs"
	"github.com/rayfiyo/kotatuneko-rest/internal/domain/entity"
)

// hello godoc
// @Summary Hello World !
// @ID      HelloWorldIndex
// @Tags    HelloWorld
// @Produce json
// @Success     200     {string}        string  "OK"
// @Router   / [get]
func Hello(c echo.Context) error {
	return c.JSON(http.StatusOK, &entity.Response{
		Int64:  1,
		String: "example",
		World: &entity.Item{
			Text: "hello world !",
		},
	})
}
