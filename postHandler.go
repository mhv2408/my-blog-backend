package main

import (
	"html/template"
	"log"
	"net/http"
)

func postHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	tmpl, err := template.ParseFiles("templates/createPost.html")

	if err != nil {
		log.Fatal("cannot parse the html file for post creation: ", err)
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Fatal("cannot execute the html file for post creation: ", err)
	}
}
