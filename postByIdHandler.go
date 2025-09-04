package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) postByIdHandler(w http.ResponseWriter, r *http.Request) {
	type payload struct {
		Id uuid.UUID `json:"id"`
	}
	decoder := json.NewDecoder(r.Body)

	var data payload

	err := decoder.Decode(&data)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "cannot decode the json for post id", err)
		return
	}

	dbPost, err := cfg.db.GetPostByID(r.Context(), data.Id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "cannot find the post with the id", err)
		return
	}
	respondWithJson(w, http.StatusOK, Post{
		Id:      dbPost.PostsID,
		Title:   dbPost.Title,
		Summary: dbPost.Summary,
		Post:    dbPost.Post,
	})
}
