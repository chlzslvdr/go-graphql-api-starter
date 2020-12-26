package models

type Lookups struct {
	ID string `json:"id"`
}

// IsEntity if the lookups is an entity
func (Lookups) IsEntity() {}
