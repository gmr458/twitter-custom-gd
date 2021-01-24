package middlewares

import (
	"net/http"

	"github.com/gmr458/twitter-custom-gd/database"
)

// CheckDB : Middleware that checks the connection to the database
func CheckDB(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		if database.CheckConnection() == 0 {
			http.Error(w, "Lost connection to the database", 500)
			return
		}

		next.ServeHTTP(w, r)
	}

}
