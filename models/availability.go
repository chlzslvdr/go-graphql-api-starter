package models

import "time"

type Availability struct {
	ID           *int       `json:"id" gorm:"column:availability_id;"`
	Availability *string    `json:"availability" gorm:"column:availability;"`
	IsActive     bool       `json:"isActive" gorm:"column:is_active;"`
	CreatedWhen  *time.Time `json:"createdWhen" gorm:"column:created_when;"`
	UpdatedWhen  *time.Time `json:"updatedWhen" gorm:"column:updated_when;"`
}
