package services

import domain_entities "go-hexa/internal/core/domain/entities"

type IUserService interface {
	FindAll() *[]domain_entities.UserEntity
}
