package internal

import (
	"go-hexa/internal/core/services"
	"go-hexa/internal/handlers"
	"go-hexa/internal/repositories"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)



func Bootstrap(app *fiber.App, conf *viper.Viper)  {
	userInit(app)
}

/*
*users init
*/
func userInit(app *fiber.App)  {
	userRepository := repositories.NewUserRepository()
	userservice := services.NewUserServie(userRepository)
	userHandler := handlers.NewUserHandler(userservice,app)

	userHandler.InitRouter()
}