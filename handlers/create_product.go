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
		util.ErrorResponse(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	products := db.GetProductList()

	status, messsage := validateProduct(newProduct, products, false)
	if !status {
		util.ErrorResponse(w, messsage, http.StatusUnprocessableEntity)
		return
	}

	newProduct.ID = len(products) + 1

	db.StoreProduct(newProduct)

	util.SuccessResponse(w, newProduct, http.StatusCreated)
}

func validateProduct(product db.Product, products []db.Product, isUpdate bool) (bool, string) {
	for i := 0; i < len(products); i++ {
		if product.Title == "" {
			return false, "Title required."
		}

		if !isUpdate {
			if products[i].Title == product.Title {
				return false, "Already exists"
			}
		} else {
			if products[i].Title == product.Title && products[i].ID == product.ID {
				return false, "Already exists"
			}
		}
	}

	return true, ""
}
