package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

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
		log.Fatal("you are not authorized")
	}
	fmt.Println("hello " + user_name)

}
