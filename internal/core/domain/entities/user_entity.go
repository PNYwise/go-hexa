package entities

import (
	"go-hexa/internal/core/domain/enums/role"
	"go-hexa/internal/core/domain/enums/status"
)




type UserEntity struct {
	Id  int64 `json:"id"`
	Name  string `json:"name"`
	Address string `json:"address"`
	Status status.Enum `json:"status"`
	Role role.Enum `json:"role"`
}