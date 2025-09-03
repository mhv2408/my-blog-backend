package main

import (
	"net/http"
)

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	respondWithJson(w, http.StatusOK, serveJson())
}
