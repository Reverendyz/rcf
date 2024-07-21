package handlers

import (
	"net/http"
	"reverendyz/rcf/internal/types"
	"reverendyz/rcf/pkg/participant"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SaveParticipant(c *gin.Context) {
	var p = &types.Participant{}
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.S().Errorf("%v", err)
		return
	}
	if err := participant.SaveParticipant(p); err != nil {
		zap.S().Errorf("%v", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "The participant could not be created"})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Participant created successfully"})
}

func ListParticipants(c *gin.Context) {
	participants, err := participant.ListParticipants()
	if err != nil {
		zap.S().Error("%v", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Status": "error", "message": err})
		return
	}
	c.IndentedJSON(http.StatusAccepted, gin.H{"participants": participants})
}
