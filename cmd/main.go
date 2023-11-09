package main

import (
	"go-hexa/internal"
	"go-hexa/internal/core/configs"
	"log"

	"github.com/gofiber/fiber/v2"
)



func main() {
	configs.ConnectDb()
	app := fiber.New()
	conf := configs.New()

	port := conf.GetString("app.port")
	internal.Bootstrap(app,conf)

	log.Fatal(app.Listen(":"+port))
}