package models

import "time"

type Gender struct {
	ID          *int       `json:"id" gorm:"column:gender_id;"`
	Gender      *string    `json:"gender" gorm:"column:gender;"`
	IsActive    bool       `json:"isActive" gorm:"column:is_active;"`
	CreatedWhen *time.Time `json:"createdWhen" gorm:"column:created_when;"`
	UpdatedWhen *time.Time `json:"updatedWhen" gorm:"column:updated_when;"`
}
