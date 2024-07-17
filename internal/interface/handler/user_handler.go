package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/rayfiyo/kotatuneko-rest/gen/user/entity"
	"github.com/rayfiyo/kotatuneko-rest/internal/usecase"
	"github.com/rayfiyo/kotatuneko-rest/pkg/errors"
)

type UserHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) *UserHandler {
	return &UserHandler{userUsecase: userUsecase}
}

func (h *UserHandler) RegisterUser(c echo.Context) error {
	var req struct {
		UserName string `json:"user_name"`
		Password string `json:"password"`
	}
	if err := c.Bind(&req); err != nil {
		log.Printf("%v: %v \n", errors.ErrInvalidRequest, err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, err := h.userUsecase.RegisterUser(c.Request().Context(), req.UserName, req.Password)
	if err != nil {
		log.Printf("%v: %v \n", errors.ErrInternalServerError, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) LoginUser(c echo.Context) error {
	var req struct {
		UserName string `json:"user_name"`
		Password string `json:"password"`
	}
	if err := c.Bind(&req); err != nil {
		log.Printf("%v: %v \n", errors.ErrInvalidRequest, err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, err := h.userUsecase.LoginUser(c.Request().Context(), req.UserName, req.Password)
	if err != nil {
		log.Printf("%v: %v \n", errors.ErrInternalServerError, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) UpdateUser(c echo.Context) error {
	var user entity.User
	if err := c.Bind(&user); err != nil {
		log.Printf("%v: %v \n", errors.ErrInvalidRequest, err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if user.Id == "" {
		// [TODO] user id がないとき探してくる
		log.Println(user.Id)
	}

	if err := h.userUsecase.UpdateUser(c.Request().Context(), &user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

func (h *UserHandler) GetRanking(c echo.Context) error {
	limitStr := c.QueryParam("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		log.Printf("%v: invalid limit\n", errors.ErrInvalidRequest)
		return echo.NewHTTPError(http.StatusBadRequest, "invalid limit")
	}

	users, err := h.userUsecase.GetRanking(c.Request().Context(), limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, users)
}
