package main

import (
	"go-hexa/internal"
	"go-hexa/internal/core/configs"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	configs.ConnectDb()
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
