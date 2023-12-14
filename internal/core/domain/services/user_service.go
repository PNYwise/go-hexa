package services

import (
	"go-hexa/internal/core/domain/models/requests"
	"go-hexa/internal/core/domain/models/responses"
)

type IUserService interface {
	FindAll(paginationRequest *requests.PaginationRequest) (*[]responses.UserResponse, *responses.PaginationResponse)
	FindOne(id uint) (*responses.UserResponse, error)
	Create(request *requests.UserRequest) (*responses.UserResponse, error)
}
