package handler

import (
	"Api1/modle"
	"encoding/json"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser modle.User
	json.NewDecoder(r.Body).Decode(&newUser)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}
