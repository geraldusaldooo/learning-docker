package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v3"
	_ "github.com/lib/pq"
)

func main() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Gagal terhubung ke database: %v\n", err)
	}
	defer db.Close()

	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {

		_, err := http.Get("http://localhost:4331")

		if err != nil {
			log.Println("Ping failed", err)
			return c.SendString("Connection Failed")
		}
		log.Printf("Ping succesful")
		return c.SendString("Pong!")
	})

	app.Get("/ping", func(c fiber.Ctx) error {

		_, err = db.Exec("INSERT INTO timestamp (timestamp) VALUES ($1)", time.Now())
		if err != nil {
			log.Printf("Ping failed :%v\n", err)
			return c.SendString("Ping failed")
		}

		log.Printf("Ping successful")
		return c.SendString("Pong")
	})

	log.Fatal(app.Listen(":77"))
}
