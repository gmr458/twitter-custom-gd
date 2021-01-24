package middlewares

import (
	"net/http"

	"github.com/gmr458/twitter-custom-gd/routers"
)

// ValidJWT : ValidJWT
func ValidJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		_, _, _, err := routers.ProcessToken(r.Header.Get("Authorization"))

		if err != nil {
			http.Error(w, "Error token! "+err.Error(), http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)

	}
}
