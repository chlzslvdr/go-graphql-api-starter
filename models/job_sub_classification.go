package models

import "time"

type JobSubClassification struct {
	ID                  *int       `json:"id" gorm:"column:job_sub_classification_id;"`
	Name                *string    `json:"name" gorm:"column:job_sub_classification;"`
	JobClassificationID *int       `json:"jobClassificationId" gorm:"column:job_classification_id;"`
	IsActive            bool       `json:"isActive" gorm:"column:is_active;"`
	CreatedWhen         *time.Time `json:"createdWhen" gorm:"column:created_when;"`
	UpdatedWhen         *time.Time `json:"updatedWhen" gorm:"column:updated_when;"`
}
