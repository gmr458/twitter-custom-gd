package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/gmr458/twitter-custom-gd/models"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InsertRecord : Insert a record into the database
func InsertRecord(u models.User) (string, bool, error) {

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

	db := MongoConn.Database("twitter-custom-db")
	col := db.Collection("users")

	u.Password, _ = EncryptPassword(u.Password)

	result, err := col.InsertOne(ctx, u)

	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)

	return ObjID.String(), true, nil
}
