package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/gmr458/twitter-custom-gd/models"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
)

// CheckTheUserAlreadyExists : Check the user already exists
func CheckTheUserAlreadyExists(email string) (models.User, bool, string) {

	var err error = godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var dbName string = os.Getenv("DB_NAME")

	if dbName == "" {
		log.Fatal("Error DB_NAME, DB_NAME = \"\"")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConn.Database(dbName)
	col := db.Collection("users")

	condition := bson.M{"email": email}

	var result models.User

	err = col.FindOne(ctx, condition).Decode(&result)

	ID := result.ID.Hex()

	if err != nil {
		return result, false, ID
	}

	return result, true, ID
}
