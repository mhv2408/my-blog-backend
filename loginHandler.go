package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/mhv2408/my-blog/internal/auth"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		log.Fatal("unable to parse the form: ", err)
	}

	err = godotenv.Load()
	if err != nil {
		log.Fatal("cannot load the env file: ", err)
	}
	user_name, password := os.Getenv("BLOG_USERNAME"), os.Getenv("BLOG_PASSWORD")

	if r.FormValue("username") != user_name || password != auth.HashPassword(r.FormValue("password")) {
		//log.Fatal("you are not authorized")
		http.ServeFile(w, r, "templates/unauthorized.html")
		return
	}
	// now you login
	expiration := time.Now().Add(30 * time.Minute) // 30 min expiration time
	cookie := http.Cookie{
		Name:    "login",
		Value:   "login-cookie",
		Expires: expiration,
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/editor/post", http.StatusFound)

}
