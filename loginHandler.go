package main

import (
	"fmt"
	"log"
	"net/http"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		log.Fatal("unable to parse the form: ", err)
	}

	fmt.Println(r.FormValue("username"), r.FormValue(("password")))
}
