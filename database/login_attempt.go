package database

import (
	"github.com/gmr458/twitter-custom-gd/models"
	"golang.org/x/crypto/bcrypt"
)

// LoginAttempt : LoginAttempt
func LoginAttempt(email string, password string) (models.User, bool) {

	usu, found, _ := CheckTheUserAlreadyExists(email)

	if found == false {
		return usu, false
	}

	passwordBytes := []byte(password)
	passwordDB := []byte(usu.Password)
	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)

	if err != nil {
		return usu, false
	}

	return usu, true
}
