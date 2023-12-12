package requests

import (
	"go-hexa/internal/core/domain/enums/role"
	"go-hexa/internal/core/domain/enums/status"
)

type UserRequest struct {
	Name     string      `json:"name" validate:"required,alphanum,min=3"`
	Email    string      `json:"email" validate:"required,email"`
	Password string      `json:"password" validate:"required,alphanum"`
	Address  string      `json:"address" validate:"alphanum"`
	Status   status.Enum `json:"status" validate:"required,alpha"`
	Role     role.Enum   `json:"role" validate:"required,alpha"`
}
