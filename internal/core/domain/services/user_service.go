package services

import (
	"go-hexa/internal/core/domain/entities"
	"go-hexa/internal/core/domain/models/requests"
	"go-hexa/internal/core/domain/models/responses"

	"github.com/gofiber/fiber/v2"
)

type IUserService interface {
	FindAll(paginationRequest *requests.PaginationRequest) (*[]entities.UserEntity, *responses.PaginationResponse)
	FindOne(id uint) (*entities.UserEntity, *fiber.Error)
}
