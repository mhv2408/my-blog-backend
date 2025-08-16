package main

import (
	"html/template"
	"log"
	"net/http"
)

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatal("cannot parse the html file: ", err)
	}
	err = tmpl.Execute(w, "")
	if err != nil {
		log.Fatal("cannot execute the template: ", err)
	}
}
