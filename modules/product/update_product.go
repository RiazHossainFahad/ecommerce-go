package product

import (
	"ecommerce/db"
	"ecommerce/util"
	"encoding/json"
	"net/http"
	"strconv"
)

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	requestId := r.PathValue("id")

	id, err := strconv.Atoi(requestId)
	if err != nil {
		util.ErrorResponse(w, "Invalid ID given.", http.StatusBadRequest)
		return
	}

	var product db.Product

	jsonErr := json.NewDecoder(r.Body).Decode(&product)
	if jsonErr != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	products := db.GetProductList()

	status, messsage := validateProduct(product, products, true)
	if !status {
		util.ErrorResponse(w, messsage, http.StatusUnprocessableEntity)
		return
	}

	product.ID = id

	status, messsage = db.UpdateProduct(id, product)
	if !status {
		util.ErrorResponse(w, messsage, http.StatusUnprocessableEntity)
		return
	}

	util.SuccessResponse(w, map[string]any{"status": true, "messsage": messsage}, http.StatusAccepted)
}
