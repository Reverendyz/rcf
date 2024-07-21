package router

import (
	"reverendyz/rcf/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/participants", handlers.ListParticipants)
	r.POST("/participant/add", handlers.SaveParticipant)
	r.GET("/expenses", handlers.ListExpenses)
	r.POST("/expenses/add", handlers.SaveExpense)
	r.PUT("/expenses/bind/:expenseID/:participantID", handlers.BindParticipantToExpense)

	return r
}
