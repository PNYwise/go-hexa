package main

import (
	"go-hexa/internal"
	"go-hexa/internal/core/configs"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// set time to utc
	time.Local = time.UTC

	configs.ConnectDb()
	configs.App()
	defer func() {
		if err := configs.CloseDb(); err != nil {
			log.Fatalf("Error closing database connection: %v", err)
		}
	}()

	app := fiber.New()
	conf := configs.New()

	port := conf.GetString("app.port")
	internal.Bootstrap(app, configs.DB.Db, conf)

	log.Fatal(app.Listen(":" + port))
}
