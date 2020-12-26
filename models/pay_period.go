package models

import "time"

type PayPeriod struct {
	ID          *int       `json:"id" gorm:"column:pay_period_id;"`
	Period      *string    `json:"period" gorm:"column:pay_period;"`
	IsActive    bool       `json:"isActive" gorm:"column:is_active;"`
	CreatedWhen *time.Time `json:"createdWhen" gorm:"column:created_when;"`
	UpdatedWhen *time.Time `json:"updatedWhen" gorm:"column:updated_when;"`
}
