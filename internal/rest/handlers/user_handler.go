package handlers

import (
	"go-hexa/internal/core/domain/models/requests"
	"go-hexa/internal/core/domain/services"

	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	app        *fiber.App
	userServie services.IUserService
}

func NewUserHandler(
	userServie services.IUserService,
	app *fiber.App,
) *userHandler {
	return &userHandler{
		app:        app,
		userServie: userServie,
	}
}

func (u *userHandler) InitRouter() {
	api := u.app.Group("users")
	api.Get("list", u.FindAll)
	api.Get("detail/:id", u.FindOne)
}

func (u *userHandler) FindAll(ctx *fiber.Ctx) error {
	paginationRequest := &requests.PaginationRequest{}
	if err := ctx.QueryParser(paginationRequest); err != nil {
		panic(err)
	}
	users, pagination := u.userServie.FindAll(paginationRequest)
	response := NewApiResponseList(200, users, pagination)

	return ctx.JSON(response)
}

func (u *userHandler) FindOne(ctx *fiber.Ctx) error {
	param := struct {
		ID uint `params:"id"`
	}{}
	if err := ctx.ParamsParser(&param); err != nil {
		panic(err)
	}
	users, err := u.userServie.FindOne(param.ID)
	if err != nil {
		return err
	}
	response := NewApiResponseDetail(200, users)
	return ctx.JSON(response)
}
