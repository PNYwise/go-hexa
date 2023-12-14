package responses

import (
	"go-hexa/internal/core/domain/enums/role"
	"go-hexa/internal/core/domain/enums/status"
)

type UserResponse struct {
	ID      uint        `json:"id"`
	Name    string      `json:"name"`
	Email   string      `json:"email"`
	Address string      `json:"address"`
	Status  status.Enum `json:"status"`
	Role    role.Enum   `json:"role"`
}
