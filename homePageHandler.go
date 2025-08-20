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
	blog_data := serveJson()
	for _, blog := range blog_data {
		//fmt.Println(blog)
		data := struct {
			Title   string
			Summary string
			Post    string
		}{
			Title:   blog.Title,
			Summary: blog.Summary,
			Post:    blog.Post,
		}
		//fmt.Print(data)
		err = tmpl.Execute(w, data)
	}

	if err != nil {
		log.Fatal("cannot execute the template: ", err)
	}
}
