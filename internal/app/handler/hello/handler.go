package hello

import (
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/rayfiyo/kotatuneko-rest/docs"
	// "github.com/rayfiyo/kotatuneko-rest/internal/domain/entity"
)

type Response struct {
	// "github.com/rayfiyo/kotatuneko-rest/internal/domain/entity"
	Int64  int64  `json:"int64"`
	String string `json:"string"`
	World  *Item  `json:"world"`
}

type Item struct {
	// "github.com/rayfiyo/kotatuneko-rest/internal/domain/entity"
	Text string `json:"text"`
}

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
	return c.JSON(http.StatusOK, &Response{
		Int64:  1,
		String: "example",
		World: &Item{
			Text: "hello world !",
		},
	})
}
