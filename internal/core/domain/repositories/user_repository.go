package repositories

import (
	"go-hexa/internal/core/domain/entities"
	"go-hexa/internal/core/domain/models/requests"
	"go-hexa/internal/core/domain/models/responses"
)

type IUserRepository interface {
	FindAll(paginationRequest *requests.PaginationRequest) (*[]entities.UserEntity, *responses.PaginationResponse)
	FindOne(id uint) (*entities.UserEntity, error)
	FindOneByEmail(email string) (*entities.UserEntity, error)
	Create(user *entities.UserEntity) (*entities.UserEntity, error)
}
