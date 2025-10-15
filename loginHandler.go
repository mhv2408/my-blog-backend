package main

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/mhv2408/my-blog/internal/auth"
)

type LoginDetails struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var data LoginDetails

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "unable to verify login creds", err)
		return
	}

	if os.Getenv("DOCKER_ENV") == "" { // running locally
		if err := godotenv.Load(); err != nil {
			respondWithError(w, http.StatusInternalServerError, "unable to verify login creds", err)
			return
		}
	}
	user_name, password := os.Getenv("BLOG_USERNAME"), os.Getenv("BLOG_PASSWORD")

	if data.Username != user_name || password != auth.HashPassword(data.Password) {
		//log.Fatal("you are not authorized")
		respondWithError(w, http.StatusUnauthorized, "Incorrect creds", nil)
		return
	}
	// now you login
	expiration := time.Now().Add(30 * time.Minute) // 30 min expiration time
	cookie := http.Cookie{
		Name:     "login",
		Value:    "login-cookie",
		Expires:  expiration,
		Path:     "/",                   // makes cookie valid for the whole site
		HttpOnly: true,                  // prevent JS from reading cookie (good for security)
		Secure:   true,                  // set true in production with HTTPS
		SameSite: http.SameSiteNoneMode, // allow cross-site with https
	}
	http.SetCookie(w, &cookie)
	respondWithJson(w, http.StatusOK, nil)

}
