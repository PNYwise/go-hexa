package main

import (
	"go-hexa/internal"
	"go-hexa/internal/core/configs"
	"go-hexa/internal/rest/handlers"
	"log"
	"time"

	_ "go-hexa/docs/api"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// @title GO-HEXA App
// @version 1.0
// @description This is an API for GO-HEXA Application

// @contact.name PNYwise
// @contact.email dinopuguh@gmail.com

// @host localhost:8081
// @BasePath /api
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

	app := fiber.New(fiber.Config{
		ErrorHandler: handlers.ErrorHandler,
	})

	app.Use(recover.New())
	app.Use(cors.New())

	//swagger route
	app.Get("/docs/*", swagger.HandlerDefault) // default

	conf := configs.New()
	internal.Bootstrap(app, configs.DB.Db, conf)

	port := conf.GetString("app.port")
	log.Fatal(app.Listen(":" + port))
}
