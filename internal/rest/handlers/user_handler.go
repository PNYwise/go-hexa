package handlers

import (
	"go-hexa/internal/core/domain/entities"
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
	api := u.app.Group("api/users")
	api.Get("list", u.FindAll)
	api.Get("detail/:id", u.FindOne)
	api.Post("create", u.Create)
}

// @Summary Find All User.
// @Description get users.
// @Tags User
// @Accept */*
// @Produce json
// @Param page query int true "int valid" mininum(1)
// @Param take query int true "int valid" mininum(1)
// @Param order query string true "string default" default("DESC")
// @Success 200 {object} handlers.ApiResponseList[[]entities.UserEntity]
// @Router /users/list [get]
func (u *userHandler) FindAll(ctx *fiber.Ctx) error {
	paginationRequest := &requests.PaginationRequest{}
	if err := ctx.QueryParser(paginationRequest); err != nil {
		panic(err)
	}
	users, pagination := u.userServie.FindAll(paginationRequest)
	response := NewApiResponseList[*[]entities.UserEntity](200, users, pagination)

	return ctx.JSON(response)
}

// @Summary Find One User.
// @Description get user by id.
// @Tags User
// @Accept */*
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} handlers.ApiResponseDetail[entities.UserEntity]
// @Router /users/detail/{id} [get]
func (u *userHandler) FindOne(ctx *fiber.Ctx) error {
	param := struct {
		ID uint `params:"id"`
	}{}
	if err := ctx.ParamsParser(&param); err != nil {
		panic(err)
	}
	users, err := u.userServie.FindOne(param.ID)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	response := NewApiResponseDetail[*entities.UserEntity](200, users)
	return ctx.JSON(response)
}

// @Summary Create User.
// @Description cteate user.
// @Tags User
// @Accept Application/Json
// @Produce json
// @Param request body requests.UserRequest true "body"
// @Success 200 {object} handlers.ApiResponseDetail[entities.UserEntity]
// @Router /users/create [post]
func (u *userHandler) Create(ctx *fiber.Ctx) error {
	request := &requests.UserRequest{}
	if err := ctx.BodyParser(&request); err != nil {
		panic(err)
	}
	user, err := u.userServie.Create(request)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	// Validation

	response := NewApiResponseDetail[*entities.UserEntity](200, user)
	return ctx.JSON(response)
}
