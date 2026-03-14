package handlers

import (
	"ecommerce/db"
	"ecommerce/util"
	"encoding/json"
	"net/http"
)

func StoreUser(w http.ResponseWriter, r *http.Request) {
	var newUser db.User

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		util.ErrorResponse(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	users := db.GetUserList()

	status, messsage := validateUser(newUser, users, false)
	if !status {
		util.ErrorResponse(w, messsage, http.StatusUnprocessableEntity)
		return
	}

	newUser.ID = len(users) + 1

	db.StoreUser(newUser)

	util.SuccessResponse(w, newUser, http.StatusCreated)
}

func validateUser(user db.User, users []db.User, isUpdate bool) (bool, string) {
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
			if users[i].Email == user.Email && users[i].ID == user.ID {
				return false, "Already exists"
			}
		}
	}

	return true, ""
}
