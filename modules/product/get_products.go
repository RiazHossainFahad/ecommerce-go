package product

import (
	"ecommerce/util"
	"net/http"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {

	list, err := h.productService.List()

	if err != nil {
		util.ErrorResponse(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if len(list) <= 0 {
		util.ErrorResponse(w, "Empty products", http.StatusNotFound)
		return
	}

	util.SuccessResponse(w, list, http.StatusOK)
}
