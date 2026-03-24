package product

import (
	"ecommerce/util"
	"net/http"
	"strconv"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {

	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = 50
	}

	list, err := h.productService.List(page, limit)

	if err != nil {
		util.ErrorResponse(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	count, err := h.productService.Count()
	if err != nil {
		util.ErrorResponse(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// if len(list) <= 0 {
	// 	util.ErrorResponse(w, "Empty products", http.StatusNotFound)
	// 	return
	// }

	util.SendPaginatedData(w, r, list, page, limit, count)
}
