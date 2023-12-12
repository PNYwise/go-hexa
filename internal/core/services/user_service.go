package services

import (
	"errors"
	"fmt"
	"go-hexa/internal/core/domain/entities"
	"go-hexa/internal/core/domain/models/requests"
	"go-hexa/internal/core/domain/models/responses"
	"go-hexa/internal/core/domain/repositories"
	"go-hexa/internal/core/domain/services"
	"go-hexa/internal/core/utils"
	"strings"
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
func (*userService) Create(request *requests.UserRequest) (*entities.UserEntity, error) {
	// Validation
	if errs := utils.Validate(request); len(errs) > 0 && errs[0].Error {
		errMsgs := make([]string, 0)

		for _, err := range errs {
			errMsgs = append(errMsgs, fmt.Sprintf(
				"[%s]: '%v' | Needs to implement '%s'",
				err.FailedField,
				err.Value,
				err.Tag,
			))
		}

		return nil, errors.New(strings.Join(errMsgs, " and "))
	}
	return nil, nil
}
