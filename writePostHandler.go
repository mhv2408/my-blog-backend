package main

import (
	"net/http"
)

func writePostHandler(w http.ResponseWriter, r *http.Request) {
	respondWithJson(w, http.StatusOK, nil)
}
