package product

import (
	"ecommerce/db"
	"ecommerce/util"
	"net/http"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	util.SuccessResponse(w, db.GetProductList(), http.StatusOK)
}
