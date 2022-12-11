package api

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/saitooooooo/sqlx-xo-example/src/models"
)

func NewRouter(e *echo.Echo, db *sqlx.DB) {
	e.GET("/hello", hello)

	e.GET("/user", getUserList(db))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World!!")
}

func getUserList(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var user []models.User
		if err := db.Select(&user, "select * from users"); err != nil {
			return err
		}

		return c.JSON(http.StatusOK, user)
	}
}
