package product

import (
	"ecommerce/util"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	requestId := r.PathValue("id")

	id, err := strconv.Atoi(requestId)
	if err != nil {
		util.ErrorResponse(w, "Invalid ID given.", http.StatusBadRequest)
		return
	}

	err = h.productRepo.Delete(id)
	if err != nil {
		util.ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.SuccessResponse(w, nil, http.StatusNoContent)
}
