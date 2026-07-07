package main

import (
	"net/http"
)

func (cfg *apiConfig) handlerReturnChirps(w http.ResponseWriter, r *http.Request) {
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
