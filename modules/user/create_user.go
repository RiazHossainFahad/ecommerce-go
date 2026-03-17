package user

import (
	"ecommerce/repo"
	"ecommerce/util"
	"encoding/json"
	"net/http"
)

func (h *Handler) StoreUser(w http.ResponseWriter, r *http.Request) {
	newUser := h.userRepo.EmptyUser()

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		util.ErrorResponse(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	users, err := h.userRepo.List()
	if err != nil {
		util.ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	status, messsage := validateUser(newUser, users, false)
	if !status {
		util.ErrorResponse(w, messsage, http.StatusUnprocessableEntity)
		return
	}

	user, err := h.userRepo.Store(newUser)
	if err != nil {
		util.ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	util.SuccessResponse(w, user, http.StatusCreated)
}

func validateUser(user repo.User, users []*repo.User, isUpdate bool) (bool, string) {
	for i := 0; i < len(users); i++ {
		if user.Email == "" {
			return false, "Email is required."
		}

		if user.Password == "" {
			return false, "Password is required."
		}

		if !isUpdate {
			if users[i].Email == user.Email {
				return false, "Already exists"
			}
		} else {
			if users[i].Email == user.Email && users[i].ID != user.ID {
				return false, "Already exists"
			}
		}
	}

	return true, ""
}
