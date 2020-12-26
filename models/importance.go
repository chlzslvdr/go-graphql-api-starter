package models

import "time"

type Importance struct {
	ID          *int       `json:"id" gorm:"column:importance_id;"`
	Importance  *string    `json:"importance" gorm:"column:importance;"`
	IsActive    bool       `json:"isActive" gorm:"column:is_active;"`
	CreatedWhen *time.Time `json:"createdWhen" gorm:"column:created_when;"`
	UpdatedWhen *time.Time `json:"updatedWhen" gorm:"column:updated_when;"`
}
