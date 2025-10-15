package main

import (
	"net/http"
)

func writeBlogHandler(w http.ResponseWriter, r *http.Request) {
	respondWithJson(w, http.StatusOK, nil)
}
