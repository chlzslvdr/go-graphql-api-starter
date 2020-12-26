package models

type CVTemplate struct {
	ID       *int    `json:"id" gorm:"column:cv_template_id;"`
	Name     *string `json:"name" gorm:"column:cv_template;"`
	URL      *string `json:"url;" gorm:"column:cv_template_url;"`
	IsActive bool    `json:"isActive" gorm:"column:is_active;"`
}
