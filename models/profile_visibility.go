package models

type ProfileVisibility struct {
	ID       *int    `json:"id" gorm:"column:profile_visibility_id;"`
	Name     *string `json:"name" gorm:"column:profile_visibility;"`
	IsActive bool    `json:"isActive" gorm:"column:is_active;"`
}
