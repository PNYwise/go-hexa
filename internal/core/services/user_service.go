package services

import (
	"fmt"
	"go-hexa/internal/core/domain/entities"
	"go-hexa/internal/core/domain/models/requests"
	"go-hexa/internal/core/domain/models/responses"
	"go-hexa/internal/core/domain/repositories"
	"go-hexa/internal/core/domain/services"

	"github.com/gofiber/fiber/v2"
)

type userService struct {
	userRepo repositories.IUserRepository
}

func NewUserServie(userRepo repositories.IUserRepository) services.IUserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (u *userService) FindAll(paginationRequest *requests.PaginationRequest) (*[]entities.UserEntity, *responses.PaginationResponse) {
	return u.userRepo.FindAll(paginationRequest)
}

func (u *userService) FindOne(id uint) (*entities.UserEntity, *fiber.Error) {
	user, err := u.userRepo.FindOne(id)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, fmt.Sprintf("user %d not found", id))
	}
	return user, nil
}
