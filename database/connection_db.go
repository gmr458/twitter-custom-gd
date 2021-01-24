package database

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var err error = godotenv.Load()

// MongoConn : Contains the function ConnectDB
var MongoConn = ConnectDB()

var uri string = os.Getenv("URI_MONGO")

var clientOptions = options.Client().ApplyURI(uri)

// ConnectDB : Allows connecting to the database
func ConnectDB() *mongo.Client {

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if uri == "" {
		log.Fatal("Error URI, URI = \"\"")
	}

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Successful connection to the database")

	return client

}

// CheckConnection : Allows to check the connection to the database
func CheckConnection() int {

	err := MongoConn.Ping(context.TODO(), nil)

	if err != nil {
		return 0
	}

	return 1

}
