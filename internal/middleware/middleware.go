package middleware

import (
	"net/http"
	"strings"

	"github.com/fobus1289/phone-book/internal/common"
	"github.com/fobus1289/phone-book/internal/model"
	"github.com/labstack/echo/v4"
)

const (
	AUTHORIZATION = "Authorization"
	BEARER        = "Bearer "
)

var (
	StatusUnauthorized  = echo.NewHTTPError(http.StatusUnauthorized)
	ParseWithExpiredJWT = common.ParseWithExpiredJWT[model.UserPayload]
)

func Auth(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {
		authorization := c.Request().Header.Get(AUTHORIZATION)

		if !strings.HasPrefix(authorization, BEARER) {
			return StatusUnauthorized
		}

		token := authorization[7:]

		user, err := ParseWithExpiredJWT(token)

		if err != nil {
			return StatusUnauthorized
		}

		c.Set("user", user)

		return next(c)
	}
}
