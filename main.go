package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("My Blog!!")
	http.HandleFunc("/", homePageHandler)
	http.HandleFunc("POST /login", loginHandler)
	http.HandleFunc("GET /editor/post", loginMiddleware(postHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
