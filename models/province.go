package models

import "time"

type Province struct {
	ID          *int       `json:"id" gorm:"column:province_id;"`
	Name        *string    `json:"name" gorm:"column:province;"`
	RegionID    *int       `json:"regionId" gorm:"column:region_id;"`
	IsActive    bool       `json:"isActive" gorm:"column:is_active;"`
	CreatedWhen *time.Time `json:"createdWhen" gorm:"column:created_when;"`
	UpdatedWhen *time.Time `json:"updatedWhen" gorm:"column:updated_when;"`
}
