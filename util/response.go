package util

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

func SuccessResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func ErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(Error{
		Status:  false,
		Message: message,
	})
}

func SendPaginatedData(w http.ResponseWriter, r *http.Request, data interface{}, page, limit, totalItems int) {
	response := NewPaginationBuilder(r, data, page, limit, totalItems).Build()
	SuccessResponse(w, response, http.StatusOK)
}
