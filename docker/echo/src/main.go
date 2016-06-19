package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/", hello)
	e.Run(standard.New(":1323"))
}

func hello(c echo.Context) error {

	cnn, err := sql.Open("mysql", "root:root@tcp(mariadb:3306)/wordpress")
	_, err = cnn.Exec("INSERT INTO user (name) VALUES (?)", "shinofara")
	if err != nil {
		return c.String(http.StatusOK, fmt.Sprintf(err.Error()))
	}

	var name string
	if err := cnn.QueryRow("SELECT name FROM user WHERE id = 1 limit 1").Scan(&name); err != nil {
		return c.String(http.StatusOK, fmt.Sprintf(err.Error()))
	}

	return c.String(http.StatusOK, fmt.Sprintf("Hello, %s", name))
}
