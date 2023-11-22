package handlers

import (
	"fmt"
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
}

func (u *userHandler) FindAll(ctx *fiber.Ctx) error {
	paginationRequest := &requests.PaginationRequest{}
	if err := ctx.QueryParser(paginationRequest); err != nil {
		panic(err)
	}
	fmt.Println(paginationRequest.Page, paginationRequest.Take, paginationRequest.Skip())
	users, pagination := u.userServie.FindAll(paginationRequest)
	response := NewApiResponseList(200, users, pagination)
	fmt.Printf("Users: %+v\n", users)

	return ctx.JSON(response)
}
