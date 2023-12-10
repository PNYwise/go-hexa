package entities

import (
	"time"

	"gorm.io/gorm"
)

type BaseEntity struct {
	ID        uint            `json:"id" gorm:"primarykey"`
	CreatedAt time.Time       `json:"created_at" gorm:"index"`
	UpdatedAt time.Time       `json:"updated_at"`
	DeletedAt *gorm.DeletedAt `json:"-" gorm:"index"`
	CreatedBy uint            `json:"created_by"`
	UpdatedBy uint            `json:"updated_by"`
	DeletedBy uint            `json:"deleted_by"`
}
