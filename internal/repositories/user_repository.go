package repositories

import (
	"go-hexa/internal/core/domain/entities"
	"go-hexa/internal/core/domain/repositories"
)

type userRepository struct {
}

func NewUserRepository() repositories.IUserRepository {
	return &userRepository{}
}

// GetAll implements domain_repositories.IUserRepository.
func (u *userRepository) FindAll() *[]entities.UserEntity {
	users := &[]entities.UserEntity{}
	return users
}
