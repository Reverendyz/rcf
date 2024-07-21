package handlers

import (
	"net/http"

	"github.com/reverendyz/rcf/internal/types"
	"github.com/reverendyz/rcf/internal/utils"
	"github.com/reverendyz/rcf/pkg/participant"

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
	utils.HandleHandlerError(err, c)
	c.IndentedJSON(http.StatusAccepted, gin.H{"participants": participants})
}
