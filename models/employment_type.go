package models

import "time"

type EmploymentType struct {
	ID          *int       `json:"id" gorm:"column:employment_type_id;"`
	Type        *string    `json:"type" gorm:"column:employment_type;"`
	IsActive    bool       `json:"isActive" gorm:"column:is_active;"`
	CreatedWhen *time.Time `json:"createdWhen" gorm:"column:created_when;"`
	UpdatedWhen *time.Time `json:"updatedWhen" gorm:"column:updated_when;"`
}
