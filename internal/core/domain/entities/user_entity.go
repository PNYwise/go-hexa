package entities

import (
	"go-hexa/internal/core/domain/enums/role"
	"go-hexa/internal/core/domain/enums/status"
)

type UserEntity struct {
	BaseEntity
	Name     string      `json:"name" gorm:"index:idx_users;type:varchar(100);not null"`
	Email    string      `json:"email" gorm:"index:idx_users;unique;type:varchar(100);not null"`
	Password string      `json:"password" gorm:"type:varchar(200);not null"`
	Address  string      `json:"address" gorm:"type:text"`
	Status   status.Enum `json:"status" gorm:"index:idx_users;type:varchar(12);not null;default:'active'"`
	Role     role.Enum   `json:"role" gorm:"index:idx_users;type:varchar(12);not null;default:'member'"`
}

func (UserEntity) TableName() string {
	return "users"
}
