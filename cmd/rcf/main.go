package main

import (
	"github.com/reverendyz/rcf/internal/db"
	"github.com/reverendyz/rcf/internal/logger"
	"github.com/reverendyz/rcf/internal/router"
	"github.com/reverendyz/rcf/internal/types"
	"github.com/reverendyz/rcf/internal/utils"

	"go.uber.org/zap"
)

func main() {
	logger.Init()
	defer logger.Sync()

	err := db.Init()
	utils.HandleError(err)

	err = runMigrations()
	utils.HandleError(err)

	r := router.SetupRouter()
	zap.S().Info("Started Server")
	utils.HandleError(err)
	if err := r.Run(); err != nil {
		utils.HandleError(err)
	}
}

func runMigrations() error {
	db := db.GetDB()
	if err := db.AutoMigrate(&types.Expense{}, &types.Participant{}); err != nil {
		return err
	}
	return nil
}
