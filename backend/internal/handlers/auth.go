package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/reverendyz/rcf/internal/auth"
	"github.com/reverendyz/rcf/internal/db"
	"github.com/reverendyz/rcf/internal/types"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var creds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := auth.Login(creds.Username, creds.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func Register(c *gin.Context) {
	var newUser types.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		zap.S().Errorf("Error hashing password: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not hash password"})
		return
	}
	newUser.Password = string(hashedPassword)

	if err := db.GetDB().Create(&newUser).Error; err != nil {
		zap.S().Errorf("Error saving user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}
