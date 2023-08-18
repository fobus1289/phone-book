package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/fobus1289/phone-book/internal/model"
	"github.com/fobus1289/phone-book/internal/services"
	"github.com/labstack/echo/v4"
)

func NewPhoneHandler(db *sql.DB) *phoneHandler {
	return &phoneHandler{
		service: services.NewPhoneService(db),
	}
}

type phoneHandler struct {
	service *services.PhoneService
}

func (p *phoneHandler) Create(c echo.Context) error {

	var (
		phone model.PhoneDto
		ctx   = c.Request().Context()
	)

	if err := c.Bind(&phone); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	user, ok := c.Get("user").(*model.UserPayload)

	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	phone.UserId = user.UserId

	id, err := p.service.Create(ctx, phone)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"id": id,
	})
}

func (p *phoneHandler) Update(c echo.Context) error {

	var (
		phoneDto model.Phone
		ctx      = c.Request().Context()
	)

	if err := c.Bind(&phoneDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	user, ok := c.Get("user").(*model.UserPayload)

	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	phoneDto.UserId = user.UserId

	if err := p.service.Update(ctx, phoneDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	return c.NoContent(http.StatusOK)
}

func (p *phoneHandler) FindByNumber(c echo.Context) error {

	var (
		number = c.QueryParam("q")
		ctx    = c.Request().Context()
	)

	phones, err := p.service.FindByNumber(ctx, number)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	return c.JSON(http.StatusOK, phones)
}

func (p *phoneHandler) Delete(c echo.Context) error {

	var (
		phoneId int
		ctx     = c.Request().Context()
	)

	phoneId, err := strconv.Atoi(c.QueryParam("phone_id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	user, ok := c.Get("user").(*model.UserPayload)

	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	err = p.service.Delete(ctx, model.PhoneDeleteDto{
		Id:     phoneId,
		UserId: user.UserId,
	})

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	return nil
}
