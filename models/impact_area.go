package models

import "time"

type ImpactArea struct {
	ID          *int       `json:"id" gorm:"column:impact_area_id;"`
	Name        *string    `json:"name" gorm:"column:impact_area;"`
	IsActive    bool       `json:"isActive" gorm:"column:is_active;"`
	CreatedWhen *time.Time `json:"createdWhen" gorm:"column:created_when;"`
	UpdatedWhen *time.Time `json:"updatedWhen" gorm:"column:updated_when;"`
}
