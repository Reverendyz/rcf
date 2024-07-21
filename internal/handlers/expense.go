package handlers

import (
	"net/http"
	"strconv"

	"github.com/reverendyz/rcf/internal/types"
	"github.com/reverendyz/rcf/internal/utils"
	"github.com/reverendyz/rcf/pkg/expense"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SaveExpense(c *gin.Context) {
	var e = &types.Expense{}
	if err := c.ShouldBindJSON(&e); err != nil {
		zap.S().Errorf("%v", err)
		return
	}
	if err := expense.SaveExpense(e); err != nil {
		zap.S().Errorf("%v", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "The Expense could not be created"})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Expense created successfully"})
}

func ListExpenses(c *gin.Context) {
	exp, err := expense.ListExpenses()
	utils.HandleError(err)
	c.IndentedJSON(http.StatusAccepted, gin.H{"expenses": exp})
}

func BindParticipantToExpense(c *gin.Context) {
	eID, err := strconv.Atoi(c.Param("expenseID"))
	utils.HandleHandlerError(err, c)
	pID, err := strconv.Atoi(c.Param("participantID"))
	utils.HandleHandlerError(err, c)
	err = expense.BindParticipantToExpense(uint(eID), uint(pID))
	utils.HandleHandlerError(err, c)
}
