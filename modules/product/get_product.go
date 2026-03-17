package product

import (
	"ecommerce/util"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	requestId := r.PathValue("id")

	id, err := strconv.Atoi(requestId)
	if err != nil {
		util.ErrorResponse(w, "Invalid ID given.", http.StatusBadRequest)
		return
	}

	product, err := h.productRepo.Get(id)
	if err != nil {
		util.ErrorResponse(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	if product == nil {
		util.ErrorResponse(w, "Not found", http.StatusNotFound)
		return
	}

	util.SuccessResponse(w, product, http.StatusOK)

}
