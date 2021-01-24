package jwt

import (
	"log"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gmr458/twitter-custom-gd/models"
	"github.com/joho/godotenv"
)

var err error = godotenv.Load()

// GenerateJWT : Function to generate a token
func GenerateJWT(t models.User) (string, error) {

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var secreKey string = os.Getenv("SECRET_KEY_JWT")

	if secreKey == "" {
		log.Fatal("Error SECRET_KEY_JWT, SECRET_KEY_JWT = \"\"")
	}

	myKey := []byte(secreKey)

	payload := jwt.MapClaims{
		"email":     t.Email,
		"name":      t.Name,
		"lastnames": t.Lastnames,
		"birthdate": t.Birthdate,
		"biography": t.Biography,
		"location":  t.Location,
		"website":   t.Website,
		"_id":       t.ID.Hex(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenStr, err := token.SignedString(myKey)

	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil

}
