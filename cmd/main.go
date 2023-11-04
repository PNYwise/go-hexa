package main

import (
	"go-hexa/internal"
	"go-hexa/internal/core/configs"

	"github.com/gofiber/fiber/v2"
)



func main() {
	app := fiber.New()
	conf := configs.New()
	
	port := conf.GetString("app.port")
	internal.Bootstrap(app,conf)

	app.Listen(":"+port)
}