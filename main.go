package main

import (
	"log"

	"github.com/gmr458/twitter-custom-gd/database"
	"github.com/gmr458/twitter-custom-gd/handlers"
)

func main() {

	if database.CheckConnection() == 0 {
		log.Fatal("No connection to the database")
		return
	}

	handlers.Handlers()

}
