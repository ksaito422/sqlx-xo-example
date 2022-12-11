package api

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewRouter(e *echo.Echo, db *sql.DB) {
	e.GET("/hello", hello)

}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World!!")
}
