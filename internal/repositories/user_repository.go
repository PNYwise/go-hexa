package repositories

import (
	"go-hexa/internal/core/domain/entities"
	"go-hexa/internal/core/domain/enums/role"
	"go-hexa/internal/core/domain/enums/status"
	"go-hexa/internal/core/domain/repositories"
)

type userRepository struct {
}

func NewUserRepository() repositories.IUserRepository {
	return &userRepository{}
}

// GetAll implements domain_repositories.IUserRepository.
func (u *userRepository) FindAll() *[]entities.UserEntity {
	users := &[]entities.UserEntity{
		{
			Id: 1,
			Name: "Bunayya",
			Address: "Jl. Raya Utara",
			Status: status.ACTIVE,
			Role: role.ADMIN,
		},
		{
			Id: 2,
			Name: "wahyu",
			Address: "Jl. Raya Utara",
			Status: status.ACTIVE,
			Role: role.MEMBER,
		},
	} 
	return users
}
