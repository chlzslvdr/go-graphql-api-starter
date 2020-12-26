package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Base struct {
	IsActive    *bool      `json:"isActive" gorm:"column:is_active;"`
	CreatedBy   uuid.UUID  `json:"createdBy" gorm:"column:created_by;"`
	UpdatedBy   uuid.UUID  `json:"updatedBy" gorm:"column:updated_by;"`
	CreatedWhen *time.Time `json:"createdWhen" gorm:"column:created_when;"`
	UpdatedWhen *time.Time `json:"updatedWhen" gorm:"column:updated_when;"`
}
