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

	product := h.productRepo.EmptyProduct()

	jsonErr := json.NewDecoder(r.Body).Decode(&product)
	if jsonErr != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	products, err := h.productRepo.List()

	if err != nil {
		util.ErrorResponse(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	status, messsage := validateProduct(product, products, true)
	if !status {
		util.ErrorResponse(w, messsage, http.StatusUnprocessableEntity)
		return
	}

	product.ID = id

	dbProduct, err := h.productRepo.Update(id, product)
	if err != nil {
		util.ErrorResponse(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	util.SuccessResponse(w, dbProduct, http.StatusAccepted)
}
