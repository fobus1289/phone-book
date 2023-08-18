package router

import (
	"database/sql"

	"github.com/fobus1289/phone-book/internal/handlers"
	"github.com/fobus1289/phone-book/internal/middleware"
	"github.com/labstack/echo/v4"
)

func RegisterHandler(e *echo.Echo, db *sql.DB) {

	userHandler := handlers.NewUserHandler(db)

	phoneHndler := handlers.NewPhoneHandler(db)

	userGroup := e.Group("user")
	{
		userGroup.POST("/register", userHandler.Register)
		userGroup.POST("/auth", userHandler.Auth)
		userGroup.GET("/:name", userHandler.FindByName, middleware.Auth)

		userGroup.GET("/phone", phoneHndler.FindByNumber, middleware.Auth)
		userGroup.POST("/phone", phoneHndler.Create, middleware.Auth)
		userGroup.PUT("/phone", phoneHndler.Update, middleware.Auth)
		userGroup.DELETE("/phone", phoneHndler.Delete, middleware.Auth)
	}

}
