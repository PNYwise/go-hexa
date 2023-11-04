package handlers

import (
	"go-hexa/internal/core/domain/services"

	"github.com/gofiber/fiber/v2"
)



type userHandler struct {
	app *fiber.App
	userServie services.IUserService
}

func NewUserHandler(
	userServie services.IUserService,
	app *fiber.App,
	) *userHandler  {
	return &userHandler{
		app: app,
		userServie: userServie,
	}
}

func (u *userHandler) InitRouter() {
	api := u.app.Group("users")
	api.Get("list",u.FindAll)
}

func (u *userHandler) FindAll(ctx *fiber.Ctx) error {
	users := u.userServie.FindAll()

	return ctx.JSON(users)
}