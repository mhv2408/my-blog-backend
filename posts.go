package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mhv2408/my-blog/internal/database"
)

type postDetails struct {
	Title   string `json:"title"`
	Summary string `json:"summary"`
	Post    string `json:"post"`
	Status  string `json:"status"`
}

func (cfg *apiConfig) posts(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("I am in post handler")
	var data postDetails

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "unable to decode the json", err)
		return
	}
	fmt.Println(data)
	fmt.Println(data.Title)
	fmt.Println(data.Summary)
	fmt.Println(data.Status)
	fmt.Println(data.Post)
	_, err = cfg.db.CreatePost(r.Context(), database.CreatePostParams{
		Title:   data.Title,
		Summary: data.Summary,
		Post:    data.Post,
		Status:  data.Status,
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Coudn't create the post", err)
		return
	}
	respondWithJson(w, http.StatusCreated, nil)

}
