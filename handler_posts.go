package main

import (
	"net/http"

	"github.com/davidelng/rssfeedaggregator/internal/database"
)

func (apiCfg *apiConfig) handlerGetPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := apiCfg.DB.GetPostsByUserID(r.Context(), database.GetPostsByUserIDParams{
		UserID: user.ID,
		Limit:  10,
	})
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Coudln't retrieve posts")
		return
	}

	respondWithJSON(w, http.StatusOK, databasePostsToPosts(posts))
}
