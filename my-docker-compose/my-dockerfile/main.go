package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)
	fmt.Println("Connect to: ", connStr)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Connection to database failed: ", err)
	}
	if err := db.Ping(); err != nil {
		fmt.Printf("Connection to database failed (DB_HOST: %s): %s\n", dbHost, err)
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		fmt.Println("Ouch!")
		return c.HTML(http.StatusOK, "That hurts.")
	})

	e.GET("/ping", func(c echo.Context) error {
		if err := db.Ping(); err != nil {
			fmt.Printf("Connection to database failed (DB_HOST: %s): %s\n", dbHost, err)
		}
		return c.HTML(http.StatusOK, "Connected to database")
	})

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "78"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
