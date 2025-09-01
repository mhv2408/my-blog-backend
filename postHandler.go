package main

import (
	"net/http"
)

func postHandler(w http.ResponseWriter, r *http.Request) {
	respondWithJson(w, http.StatusOK, nil)
}
