package main

import "net/http"

func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	// set all the CORS permissions for the pre-flight responses
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true") //for the tokens
		if r.Method == http.MethodOptions {
			// it is just a preflight request
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next(w, r)
	}
}
