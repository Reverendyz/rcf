package expense

import (
	"errors"

	"github.com/reverendyz/rcf/internal/db"
	"github.com/reverendyz/rcf/internal/types"
)

func SaveExpense(expense *types.Expense) error {
	db := db.GetDB()
	tx := db.Begin()
	if err := tx.Create(&expense).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func ListExpenses() ([]types.Expense, error) {
	var expenses []types.Expense
	result := db.GetDB().Preload("Participants").Find(&expenses)
	return expenses, result.Error
}

func BindParticipantToExpense(expenseID uint, participantID uint) error {
	var expense types.Expense
	var participant types.Participant

	tx := db.GetDB().Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.First(&expense, expenseID).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.First(&participant, participantID).Error; err != nil {
		tx.Rollback()
		return err
	}

	for _, exp := range participant.Expenses {
		if exp.ID == expenseID {
			tx.Rollback()
			return errors.New("participant is already bound to the expense")
		}
	}

	participant.Expenses = append(participant.Expenses, &expense)

	if err := tx.Save(&participant).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
