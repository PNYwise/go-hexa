package handlers

import (
	"go-hexa/internal/core/domain/models/requests"
	"go-hexa/internal/core/domain/models/responses"
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
// @Success 200 {object} handlers.ApiResponseList[[]responses.UserResponse]
// @Router /users/list [get]
func (u *userHandler) FindAll(ctx *fiber.Ctx) error {
	paginationRequest := &requests.PaginationRequest{}
	if err := ctx.QueryParser(paginationRequest); err != nil {
		panic(err)
	}
	users, pagination, err := u.userServie.FindAll(paginationRequest)
	if err != nil {
		return err
	}
	response := NewApiResponseList[*[]responses.UserResponse](200, users, pagination)

	return ctx.JSON(response)
}

// @Summary Find One User.
// @Description get user by id.
// @Tags User
// @Accept */*
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} handlers.ApiResponseDetail[responses.UserResponse]
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
	response := NewApiResponseDetail[*responses.UserResponse](200, users)
	return ctx.JSON(response)
}

// @Summary Create User.
// @Description create user.
// @Tags User
// @Accept Application/Json
// @Produce json
// @Param request body requests.UserRequest true "body"
// @Success 201 {object} handlers.ApiResponseDetail[responses.UserResponse]
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
	response := NewApiResponseDetail[*responses.UserResponse](201, user)
	return ctx.Status(201).JSON(response)
}
