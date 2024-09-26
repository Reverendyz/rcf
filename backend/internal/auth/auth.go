package auth

import (
	"errors"

	"github.com/reverendyz/rcf/internal/db"
	"github.com/reverendyz/rcf/internal/types"
	"golang.org/x/crypto/bcrypt"
)

func Login(username, password string) (string, error) {
	var user types.User
	if err := db.GetDB().Where("username = ?", username).First(&user).Error; err != nil {
		return "", errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("incorrect password")
	}

	token, err := GenerateJWT(user)
	if err != nil {
		return "", err
	}

	return token, nil
}
