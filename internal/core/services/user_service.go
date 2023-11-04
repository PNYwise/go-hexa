package services

import (
	"go-hexa/internal/core/domain/entities"
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

func (u *userService) FindAll() *[]entities.UserEntity {
	return u.userRepo.FindAll()
}