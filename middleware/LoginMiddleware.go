package middleware

import (
	"net/http"
)

func LoginMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 1. First check the cookie
		_, err := r.Cookie("login")

		if err != nil {
			// cookie not found
			// so redirect to login
			//respondWithError(w, http.StatusUnauthorized, "not authorized to access this page", err)
			return
		}
		// cookie is set, now move next
		next(w, r)
	}
}
