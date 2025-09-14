package main

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/gosimple/slug"
	"github.com/mhv2408/my-blog/internal/database"
)

type postDetails struct {
	Title   string `json:"title"`
	Summary string `json:"summary"`
	Post    string `json:"post"`
	Status  string `json:"status"`
}

func (cfg *apiConfig) posts(w http.ResponseWriter, r *http.Request) {
	var data postDetails

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "unable to decode the json", err)
		return
	}
	_, err = cfg.db.CreatePost(r.Context(), database.CreatePostParams{
		Title:   data.Title,
		Summary: data.Summary,
		Post:    data.Post,
		Status:  data.Status,
		Slug:    slug.Make(data.Title),
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Coudn't create the post", err)
		return
	}
	respondWithJson(w, http.StatusCreated, nil)

}
func TestPosts(t *testing.T) {

}
