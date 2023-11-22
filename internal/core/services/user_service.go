package services

import (
	"go-hexa/internal/core/domain/entities"
	"go-hexa/internal/core/domain/models/requests"
	"go-hexa/internal/core/domain/models/responses"
	"go-hexa/internal/core/domain/repositories"
	"go-hexa/internal/core/domain/services"
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
