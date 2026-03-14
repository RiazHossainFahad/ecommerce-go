package handlers

import (
	"ecommerce/db"
	"ecommerce/util"
	"encoding/json"
	"net/http"
)

func StoreProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct db.Product

	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	for i := 0; i < len(db.ProductList); i++ {
		if newProduct.Title == "" {
			util.ErrorResponse(w, "Title required.", http.StatusUnprocessableEntity)
			return
		}

		if db.ProductList[i].Title == newProduct.Title {
			util.ErrorResponse(w, "Already exists", http.StatusUnprocessableEntity)
			return
		}
	}

	newProduct.ID = len(db.ProductList) + 1

	db.ProductList = append(db.ProductList, newProduct)

	util.SuccessResponse(w, newProduct, http.StatusCreated)
}
