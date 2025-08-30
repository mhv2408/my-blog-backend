package main

import (
	"fmt"
	"net/http"
)

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("I am in homePageHandler")
	respondWithJson(w, http.StatusOK, serveJson())
}
