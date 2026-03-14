package handlers

import (
	"ecommerce/db"
	"ecommerce/util"
	"net/http"
	"strconv"
)

func GetProductById(w http.ResponseWriter, r *http.Request) {
	requestId := r.PathValue("id")

	id, err := strconv.Atoi(requestId)
	if err != nil {
		util.ErrorResponse(w, "Invalid ID given.", http.StatusBadRequest)
		return
	}

	for _, product := range db.ProductList {
		if product.ID == id {
			util.SuccessResponse(w, product, http.StatusOK)
			return
		}
	}

	util.ErrorResponse(w, "Not found", http.StatusNotFound)
}
