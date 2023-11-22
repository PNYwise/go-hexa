package services

import (
	domain_entities "go-hexa/internal/core/domain/entities"
	"go-hexa/internal/core/domain/models/requests"
	"go-hexa/internal/core/domain/models/responses"
)

type IUserService interface {
	FindAll(paginationRequest *requests.PaginationRequest) (*[]domain_entities.UserEntity, *responses.PaginationResponse)
}
