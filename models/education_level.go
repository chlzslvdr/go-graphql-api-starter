package models

import "time"

type EducationLevel struct {
	ID          *int       `json:"id" gorm:"column:education_level_id;"`
	Level       *string    `json:"level" gorm:"column:education_level;"`
	IsActive    bool       `json:"isActive" gorm:"column:is_active;"`
	CreatedWhen *time.Time `json:"createdWhen" gorm:"column:created_when;"`
	UpdatedWhen *time.Time `json:"updatedWhen" gorm:"column:updated_when;"`
}
