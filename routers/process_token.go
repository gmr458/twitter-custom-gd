package routers

import (
	"errors"
	"log"
	"os"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gmr458/twitter-custom-gd/database"
	"github.com/gmr458/twitter-custom-gd/models"
	"github.com/joho/godotenv"
)

var err error = godotenv.Load()

// Email : Email
var Email string

// IDUser : IDUser
var IDUser string

// ProcessToken : ProcessToken
func ProcessToken(tk string) (*models.Claim, bool, string, error) {

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var secreKey string = os.Getenv("SECRET_KEY_JWT")

	if secreKey == "" {
		log.Fatal("Error SECRET_KEY_JWT, SECRET_KEY_JWT = \"\"")
	}

	myKey := []byte(secreKey)

	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")

	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("token format invalid")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})

	if err != nil {

		_, found, _ := database.CheckTheUserAlreadyExists(claims.Email)

		if found == true {
			Email = claims.Email
			IDUser = claims.ID.Hex()
		}

		return claims, found, IDUser, nil

	}

	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalid")
	}

	return claims, false, string(""), err

}
