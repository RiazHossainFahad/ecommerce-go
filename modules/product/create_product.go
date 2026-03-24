package product

import (
	"ecommerce/domain"
	"ecommerce/util"
	"encoding/json"
	"net/http"
)

func (h *Handler) StoreProduct(w http.ResponseWriter, r *http.Request) {
	newProduct := h.productService.EmptyProduct()

	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		util.ErrorResponse(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	products, err := h.productService.List()

	if err != nil {
		util.ErrorResponse(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	status, messsage := validateProduct(newProduct, products, false)
	if !status {
		util.ErrorResponse(w, messsage, http.StatusUnprocessableEntity)
		return
	}

	newProduct.ID = len(products) + 1

	product, err := h.productService.Store(newProduct)

	if err != nil {
		util.ErrorResponse(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	util.SuccessResponse(w, product, http.StatusCreated)
}

func validateProduct(product domain.Product, products []*domain.Product, isUpdate bool) (bool, string) {
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
