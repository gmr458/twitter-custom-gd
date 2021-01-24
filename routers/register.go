package routers

import (
	"encoding/json"
	"net/http"

	"github.com/gmr458/twitter-custom-gd/database"
	"github.com/gmr458/twitter-custom-gd/models"
)

// Register : Register users in the database
func Register(w http.ResponseWriter, r *http.Request) {

	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error in the received data "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "The user's email is required", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "The password must be at least 6 characters", 400)
		return
	}

	_, found, _ := database.CheckTheUserAlreadyExists(t.Email)
	if found == true {
		http.Error(w, "There is already a registered user with this email", 400)
		return
	}

	_, status, err := database.InsertRecord(t)
	if err != nil {
		http.Error(w, "Error trying to register a user "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "Error trying to register a user", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
