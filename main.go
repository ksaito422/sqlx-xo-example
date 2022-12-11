package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"

	"github.com/saitooooooo/sqlx-xo-example/src/api"
)

var (
	dbUser     = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASS")
	dbDatabase = os.Getenv("DB_NAME")
	dbHost     = os.Getenv("DB_HOST")
	dbConn     = fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbDatabase)
)

func main() {
	// DBコネクション
	db, err := sqlx.Connect("mysql", dbConn)
	if err != nil {
		log.Println("fail to connect DB")
		return
	}

	// インスタンスを作成
	e := echo.New()

	api.NewRouter(e, db)

	e.Logger.Fatal(e.Start(":8000"))
}
