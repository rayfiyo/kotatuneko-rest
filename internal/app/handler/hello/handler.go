package hello

import (
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/rayfiyo/kotatuneko-rest/docs"
)

// @Summary		Hello
// @Description	Hello World 用
// @Tags		Hello
// @Accept		string
// @Produce		string

// @Success	200	{string}	string	"OK"
// @Failure	400	{string}	string	"Bad Request"
// @Failure	404	{string}	string	"Not Found"

// @Router		/hello [get]

func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
