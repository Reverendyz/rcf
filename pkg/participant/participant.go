package participant

import (
	"github.com/reverendyz/rcf/internal/db"
	"github.com/reverendyz/rcf/internal/types"
)

func SaveParticipant(participant *types.Participant) error {
	db := db.GetDB()
	tx := db.Begin()
	if err := tx.Create(&participant).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func ListParticipants() ([]types.Participant, error) {
	var participants = []types.Participant{}
	result := db.GetDB().Find(&participants)
	return participants, result.Error
}
