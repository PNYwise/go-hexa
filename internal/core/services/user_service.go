package services

import (
	"fmt"
	"go-hexa/internal/core/domain/entities"
	"go-hexa/internal/core/domain/enums/role"
	"go-hexa/internal/core/domain/enums/status"
	"go-hexa/internal/core/domain/models/requests"
	"go-hexa/internal/core/domain/models/responses"
	"go-hexa/internal/core/domain/repositories"
	"go-hexa/internal/core/domain/services"
	"go-hexa/internal/core/utils"
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

func (u *userService) FindOne(id uint) (*entities.UserEntity, error) {
	user, err := u.userRepo.FindOne(id)
	if err != nil {
		return nil, fmt.Errorf("user %d not found", id)
	}
	return user, nil
}

// Create implements services.IUserService.
func (u *userService) Create(request *requests.UserRequest) (*entities.UserEntity, error) {
	// Validation
	if errs := utils.Validate(request); len(errs) > 0 && errs[0].Error {
		return nil, utils.ValidationErrMsg(errs)
	}
	userEmail, _ := u.userRepo.FindOneByEmail(request.Email)
	if userEmail != nil {
		return nil, utils.ValidationErrMsg(
			[]utils.ErrorResponse{{FailedField: "email", Value: request.Email, Tag: "unique", Error: true}},
		)
	}

	//validateRole
	if err := role.Validate(request.Role); err != nil {
		return nil, utils.ValidationErrMsg(
			[]utils.ErrorResponse{{FailedField: "role", Value: request.Role, Tag: "Valid", Error: true}},
		)
	}

	//validateStatus
	if err := status.Validate(request.Status); err != nil {
		return nil, utils.ValidationErrMsg(
			[]utils.ErrorResponse{{FailedField: "status", Value: request.Status, Tag: "Valid", Error: true}},
		)
	}

	user := &entities.UserEntity{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
		Address:  request.Address,
		Status:   request.Status,
		Role:     request.Role,
	}
	u.userRepo.Create(user)
	return user, nil
}
