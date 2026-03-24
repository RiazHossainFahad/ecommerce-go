package product

import (
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

	product := h.productService.EmptyProduct()

	jsonErr := json.NewDecoder(r.Body).Decode(&product)
	if jsonErr != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	product.ID = id

	status, messsage := validateProduct(h, product)
	if !status {
		util.ErrorResponse(w, messsage, http.StatusUnprocessableEntity)
		return
	}

	dbProduct, err := h.productService.Update(id, product)
	if err != nil {
		util.ErrorResponse(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	util.SuccessResponse(w, dbProduct, http.StatusAccepted)
}
