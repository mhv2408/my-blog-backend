package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("My Blog!!")
	http.HandleFunc("/", homePageHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
