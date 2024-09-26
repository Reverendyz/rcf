package router

import (
	"github.com/gin-contrib/cors"
	"github.com/reverendyz/rcf/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"Content-Type"},
	}))
	r.POST("/login", handlers.Login) // TODO Test when do the auth part
	r.GET("/participants", handlers.ListParticipants)
	r.POST("/participant/add", handlers.SaveParticipant)
	r.DELETE("/participant/:id", handlers.DeleteParticipant)
	r.GET("/expenses", handlers.ListExpenses)
	r.POST("/expenses/add", handlers.SaveExpense)
	r.DELETE("/expenses/:id", handlers.DeleteExpense)
	r.PUT("/expenses/bind/:expenseID/:participantID", handlers.BindParticipantToExpense)

	return r
}
