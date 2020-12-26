package models

import "time"

type JobClassification struct {
	ID          *int       `json:"id" gorm:"column:job_classification_id;"`
	Name        *string    `json:"name" gorm:"column:job_classification;"`
	IsActive    bool       `json:"isActive" gorm:"column:is_active;"`
	CreatedWhen *time.Time `json:"createdWhen" gorm:"column:created_when;"`
	UpdatedWhen *time.Time `json:"updatedWhen" gorm:"column:updated_when;"`
}
