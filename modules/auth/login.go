package auth

import (
	"ecommerce/util"
	"encoding/json"
	"net/http"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest LoginRequest

	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		util.ErrorResponse(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	user, err := h.svc.Find(loginRequest.Email, loginRequest.Password)
	if user == nil {
		util.ErrorResponse(w, "Invalid credentials. Please try again.", http.StatusBadRequest)
		return
	}

	jwtSecret := h.cnf.JwtSecret
	jwtToken, err := util.CreateJwtToken(jwtSecret, util.Payload{
		Sub:         user.ID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		IsShopOwner: user.IsShopOwner,
	})
	if err != nil {
		util.ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := make(map[string]any)

	res["status"] = true
	res["user"] = user
	res["token"] = jwtToken

	util.SuccessResponse(w, res, http.StatusOK)
}
