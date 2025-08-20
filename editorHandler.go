package main

import (
	"html/template"
	"log"
	"net/http"
)

func editorHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	tmpl, err := template.ParseFiles("templates/login.html")
	if err != nil {
		log.Fatal("cannot parse the html file: ", err)
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Fatal("cannot execute the template: ", err)
	}

}
