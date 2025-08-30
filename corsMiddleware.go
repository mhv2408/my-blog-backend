package main

import "net/http"

func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	// set all the CORS permissions for the pre-flight responses
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Allowed-Access-Control-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == http.MethodOptions {
			// it is just a preflight request
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next(w, r)
	}
}
