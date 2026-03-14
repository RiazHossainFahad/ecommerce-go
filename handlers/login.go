package handlers

import (
	"ecommerce/db"
	"ecommerce/util"
	"encoding/json"
	"net/http"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest LoginRequest

	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		util.ErrorResponse(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	user := db.FindUser(loginRequest.Email, loginRequest.Password)
	if user == nil {
		util.ErrorResponse(w, "Invalid credentials. Please try again.", http.StatusBadRequest)
		return
	}

	util.SuccessResponse(w, user, http.StatusOK)
}
