package repositories

import (
	"go-hexa/internal/core/domain/entities"
	"go-hexa/internal/core/domain/repositories"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repositories.IUserRepository {
	return &userRepository{
		db: db,
	}
}

// GetAll implements domain_repositories.IUserRepository.
func (u *userRepository) FindAll() *[]entities.UserEntity {
	users := new([]entities.UserEntity)
	u.db.Find(&users)
	return users
}
