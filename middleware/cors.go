package middleware

import (
	"net/http"
)

// MIDDLEWARE
func HandleCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Allow specific origins
		// allowedOrigins := map[string]bool{
		// 	"http://localhost:3000": true,
		// 	"https://myapp.com":     true,
		// }

		// origin := r.Header.Get("Origin")
		// if allowedOrigins[origin] {
		// 	w.Header().Set("Access-Control-Allow-Origin", origin)
		// }

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Content-Type", "application/json")

		next.ServeHTTP(w, r)
	})
}
