package types

import "gorm.io/gorm"

type Expense struct {
	gorm.Model
	Value        float64        `json:"value"`
	Description  string         `json:"description"`
	Type         string         `json:"type"`
	Participants []*Participant `gorm:"many2many:participant_expense;"`
	Status       bool           `json:"status"`
	IsActive     bool           `json:"is_active"`
}
