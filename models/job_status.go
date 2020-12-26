package models

import "time"

type JobStatus struct {
	ID          *int       `json:"id" gorm:"column:job_status_id;"`
	Status      *string    `json:"status" gorm:"column:job_status;"`
	IsActive    bool       `json:"isActive" gorm:"column:is_active;"`
	CreatedWhen *time.Time `json:"createdWhen" gorm:"column:created_when;"`
	UpdatedWhen *time.Time `json:"updatedWhen" gorm:"column:updated_when;"`
}
