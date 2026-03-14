package middleware

import (
	"context"
	"ecommerce/db"
	"encoding/json"
	"net/http"
)

type contextKey string

const (
	requestBodyKey contextKey = "requestBody"
)

func ParseProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Decode JSON before next
		var data db.Product

		// Decode the JSON body
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, "Invalid JSON body", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), requestBodyKey, data)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
