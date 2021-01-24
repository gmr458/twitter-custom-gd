package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gmr458/twitter-custom-gd/database"
	"github.com/gmr458/twitter-custom-gd/jwt"
	"github.com/gmr458/twitter-custom-gd/models"
)

// Login : Login
func Login(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")

	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "User or password invalid "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Email is required", 400)
		return
	}

	document, exist := database.LoginAttempt(t.Email, t.Password)

	if exist == false {
		http.Error(w, "User or password invalid", 400)
		return
	}

	jwtKey, err := jwt.GenerateJWT(document)

	if err != nil {
		http.Error(w, "Error trying to generate the token "+err.Error(), 400)
		return
	}

	resp := models.ResponseLogin{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	expiryTime := time.Now().Add(24 * time.Hour)

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expiryTime,
	})

}
