package main

import (
	"net/http"
)

func loginMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 1. First check the cookie
		_, err := r.Cookie("login")

		if err != nil {
			// cookie not found
			// so redirect to login
			http.Redirect(w, r, "/editor", http.StatusFound)
			return
		}
		// cookie is set, now move next
		next(w, r)
	}
}
