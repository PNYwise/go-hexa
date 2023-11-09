package internal

import (
	"go-hexa/internal/core/services"
	"go-hexa/internal/handlers"
	"go-hexa/internal/repositories"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func Bootstrap(app *fiber.App, db *gorm.DB, conf *viper.Viper) {
	userInit(app, db)
}

/*
*users init
 */
func userInit(app *fiber.App, db *gorm.DB) {
	userRepository := repositories.NewUserRepository(db)
	userservice := services.NewUserServie(userRepository)
	userHandler := handlers.NewUserHandler(userservice, app)

	userHandler.InitRouter()
}
