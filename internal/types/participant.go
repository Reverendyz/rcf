package types

type Participant struct {
	ID       int
	Name     string
	Expenses []*Expense `gorm:"many2many:participant_expense;"`
}
