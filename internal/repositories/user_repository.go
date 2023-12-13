package repositories

import (
	"go-hexa/internal/core/domain/entities"
	"go-hexa/internal/core/domain/models/requests"
	"go-hexa/internal/core/domain/models/responses"
	"go-hexa/internal/core/domain/repositories"
	"go-hexa/internal/core/utils"

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
func (u *userRepository) FindAll(paginationRequest *requests.PaginationRequest) (*[]entities.UserEntity, *responses.PaginationResponse) {
	var count int64 = 0
	if err := u.db.Model(&entities.UserEntity{}).Count(&count).Error; err != nil {
		panic(err)
	}
	pagination := utils.GeneratePagination(paginationRequest, count)

	users := new([]entities.UserEntity)
	if err := u.db.Offset(paginationRequest.Skip()).Limit(paginationRequest.Take).Find(&users).Error; err != nil {
		panic(err)
	}

	return users, pagination
}

func (u *userRepository) FindOne(id uint) (*entities.UserEntity, error) {
	user := new(entities.UserEntity)
	if err := u.db.Select(
		"id",
		"name",
		"email",
		"password",
		"status",
		"role",
	).First(&user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userRepository) FindOneByEmail(email string) (*entities.UserEntity, error) {
	user := new(entities.UserEntity)
	if err := u.db.Select(
		"id",
		"email",
		"password",
	).Where(&entities.UserEntity{Email: email}).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userRepository) Create(user *entities.UserEntity) (*entities.UserEntity, error) {
	err := u.db.Transaction(func(tx *gorm.DB) error {
		tx.Create(&user)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return user, nil
}
