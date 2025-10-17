package middleware

import (
	"net/http"
)

var allowedOrigins = []string{
	"http://localhost:5173",
	"https://harsha-mirthinti.com",
}

func CorsMiddleware(next http.Handler) http.Handler {
	// set all the CORS permissions for the pre-flight responses
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		for _, allowedOrigin := range allowedOrigins {
			if origin == allowedOrigin {
				w.Header().Set("Access-Control-Allow-Origin", origin)
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
				w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
				w.Header().Set("Access-Control-Allow-Credentials", "true") //for the tokens
				if r.Method == http.MethodOptions {
					// it is just a preflight request
					w.WriteHeader(http.StatusNoContent)
					return
				}
				next.ServeHTTP(w, r) // does not return
				return
			}
		}
		// No origin match, so forbidden request
		http.Error(w, "Forbidden", http.StatusForbidden)

	})
}
