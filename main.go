package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("My Blog!!")
	http.HandleFunc("/", homePageHandler)
	http.HandleFunc("GET /editor", editorHandler)
	http.HandleFunc("POST /login", loginHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
