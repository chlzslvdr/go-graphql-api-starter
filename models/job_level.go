package models

import "time"

type JobLevel struct {
	ID          *int       `json:"id" gorm:"column:job_level_id;"`
	Level       *string    `json:"level" gorm:"column:job_level;"`
	IsActive    bool       `json:"isActive" gorm:"column:is_active;"`
	CreatedWhen *time.Time `json:"createdWhen" gorm:"column:created_when;"`
	UpdatedWhen *time.Time `json:"updatedWhen" gorm:"column:updated_when;"`
}
