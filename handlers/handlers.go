package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gmr458/twitter-custom-gd/middlewares"
	"github.com/gmr458/twitter-custom-gd/routers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

// Handlers : Manejadores
func Handlers() {

	var err error = godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var PORT string = os.Getenv("PORT")

	if PORT == "" {
		log.Fatal("Error PORT, PORT = \"\"")
	}

	router := mux.NewRouter()

	router.HandleFunc("/register", middlewares.CheckDB(routers.Register)).Methods("POST")
	router.HandleFunc("/login", middlewares.CheckDB(routers.Login)).Methods("POST")
	router.HandleFunc("/view-profile", middlewares.CheckDB(middlewares.ValidJWT(routers.ViewProfile))).Methods("GET")

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
