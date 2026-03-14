package handlers

import (
	"ecommerce/db"
	"ecommerce/util"
	"net/http"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	util.SuccessResponse(w, db.GetProductList(), http.StatusOK)
}
