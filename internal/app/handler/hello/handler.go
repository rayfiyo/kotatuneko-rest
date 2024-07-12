package hello

import (
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/rayfiyo/kotatuneko-rest/docs"
	"github.com/rayfiyo/kotatuneko-rest/internal/domain/entity"
)

// @Summary		Hello
// @Description	Hello World 用
// @Tags		Hello
// @Accept		json
// @Produce		json

// @Success	200	{string}	string	"OK"
// @Failure	400	{string}	string	"Bad Request"
// @Failure	404	{string}	string	"Not Found"

// @Router	/hello	[get]
func Hello(c echo.Context) error {
	return c.JSON(http.StatusOK, &entity.Response{
		Int64:  1,
		String: "example",
		World: &entity.Item{
			Text: "hello world !",
		},
	})
}
