package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func HandleHandlerError(err error, c *gin.Context) {
	if err != nil {
		zap.S().Errorf("%v", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Status": "error", "message": err})
	}
}

func HandleError(err error) {
	if err != nil {
		zap.S().Errorf("%v", err)
	}
}
