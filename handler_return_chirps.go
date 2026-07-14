package main

import (
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerReturnChirps(w http.ResponseWriter, r *http.Request) {
	authorQuery := r.URL.Query().Get("author_id")

	if authorQuery != "" {
		parseID, err := uuid.Parse(authorQuery)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Error parsing author_id", err)
			return
		}

		authorChirps, err := cfg.db.GetChirpsByAuthor(r.Context(), parseID)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Error getting authors chirps", err)
			return
		}

		chirps := []Chirp{}
		for _, authorChirp := range authorChirps {
			chirp := Chirp{
				ID:        authorChirp.ID,
				CreatedAt: authorChirp.CreatedAt,
				UpdatedAt: authorChirp.UpdatedAt,
				UserID:    authorChirp.UserID,
				Body:      authorChirp.Body,
			}
			chirps = append(chirps, chirp)
		}

		respondWithJSON(w, http.StatusOK, chirps)
		return
	}

	dbChirps, err := cfg.db.RetrieveChirps(r.Context())
	if err != nil {
		respondWithError(w, http.StatusBadGateway, err.Error(), err)
		return
	}

	chirps := []Chirp{}
	for _, dbChirp := range dbChirps {
		chirp := Chirp{
			ID:        dbChirp.ID,
			CreatedAt: dbChirp.CreatedAt,
			UpdatedAt: dbChirp.UpdatedAt,
			UserID:    dbChirp.UserID,
			Body:      dbChirp.Body,
		}
		chirps = append(chirps, chirp)
	}

	respondWithJSON(w, http.StatusOK, chirps)
}
