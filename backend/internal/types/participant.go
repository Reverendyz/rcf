package types

import "gorm.io/gorm"

type Participant struct {
	gorm.Model
	Name     string     `json:"name"`
	Expenses []*Expense `gorm:"many2many:participant_expense;"`
	IsActive bool       `json:"is_active"`
}
