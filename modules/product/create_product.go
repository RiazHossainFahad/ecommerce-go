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

	status, messsage := validateProduct(h, newProduct)
	if !status {
		util.ErrorResponse(w, messsage, http.StatusUnprocessableEntity)
		return
	}

	product, err := h.productService.Store(newProduct)

	if err != nil {
		util.ErrorResponse(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	util.SuccessResponse(w, product, http.StatusCreated)
}

func validateProduct(h *Handler, product domain.Product) (bool, string) {
	if product.Title == "" {
		return false, "Title required."
	}

	exists, err := h.productService.Unique(product.Title, product.ID)

	if err != nil {
		return false, err.Error()
	}

	if exists {
		return false, "Product already exists"
	}

	return true, ""
}
