package internal

import (
	"go-hexa/internal/core/services"
	grpchandlers "go-hexa/internal/grpc/handlers"
	"go-hexa/internal/repositories"
	"go-hexa/internal/rest/handlers"
	"go-hexa/proto"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func Bootstrap(app *fiber.App, db *gorm.DB, conf *viper.Viper) {
	userInit(app, db)
}
func BootstrapGrpc(srv *grpc.Server, db *gorm.DB, conf *viper.Viper) {
	userGrpcInit(srv, db)
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

func userGrpcInit(srv *grpc.Server, db *gorm.DB) {
	userRepository := repositories.NewUserRepository(db)
	userservice := services.NewUserServie(userRepository)
	userHandler := grpchandlers.NewUserHandler(userservice)

	proto.RegisterUserServer(srv, userHandler)
}
