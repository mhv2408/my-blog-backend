package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	fmt.Println("responding with json")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)

}

func respondWithError(w http.ResponseWriter, code int, msg string, err error) {
	fmt.Println("responding with error")
	if err != nil {
		log.Fatal(err)
	}
	if code > 499 {
		log.Printf("Responding with 5XX error: %s", msg)
	}
	type errorResp struct {
		Error string `json:"error"`
	}
	respondWithJson(w, code, errorResp{
		Error: msg,
	})
}
