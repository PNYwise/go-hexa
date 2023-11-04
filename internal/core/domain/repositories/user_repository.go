package repositories

import domain_entities "go-hexa/internal/core/domain/entities"


type IUserRepository interface {
	FindAll() *[]domain_entities.UserEntity
}