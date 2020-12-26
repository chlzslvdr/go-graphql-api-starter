package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Skill struct {
	ID          uuid.UUID  `json:"id" gorm:"column:skill_id;"`
	Name        *string    `json:"name" gorm:"column:skill;"`
	IsActive    bool       `json:"isActive" gorm:"column:is_active;"`
	CreatedWhen *time.Time `json:"createdWhen" gorm:"column:created_when;"`
	UpdatedWhen *time.Time `json:"updatedWhen" gorm:"column:updated_when;"`
}
