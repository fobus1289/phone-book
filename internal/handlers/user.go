package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/fobus1289/phone-book/internal/model"
	"github.com/fobus1289/phone-book/internal/services"
	"github.com/labstack/echo/v4"
)

func NewUserHandler(db *sql.DB) *userHandler {
	return &userHandler{
		service: services.NewUserService(db),
	}
}

type userHandler struct {
	service *services.UserService
}

func (u *userHandler) Register(c echo.Context) error {

	age, err := strconv.Atoi(c.FormValue("age"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	ctx := c.Request().Context()

	user := model.UserDto{
		Login:    c.FormValue("login"),
		Password: c.FormValue("password"),
		Name:     c.FormValue("name"),
		Age:      age,
	}

	id, err := u.service.Create(ctx, user)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, map[string]int64{
		"id": id,
	})
}

func (u *userHandler) FindByName(c echo.Context) error {

	var (
		name = c.Param("name")
		ctx  = c.Request().Context()
	)

	result, err := u.service.FindByName(ctx, name)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	return c.JSON(http.StatusOK, result)
}

func (u *userHandler) Auth(c echo.Context) error {

	var userSignIn model.UserSignIn

	if err := c.Bind(&userSignIn); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	ctx := c.Request().Context()

	token, err := u.service.FindByLoginAndPassword(ctx, userSignIn)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}
