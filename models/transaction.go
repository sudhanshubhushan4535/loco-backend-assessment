package models

type Transaction struct {
	ID       int64   `json:"id"`
	Amount   float64 `json:"amount" validate:"required,gt=0"`
	Type     string  `json:"type" validate:"required"`
	ParentID *int64  `json:"parent_id,omitempty" validate:"omitempty,gt=0"`
}
