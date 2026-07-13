package main

import (
	"encoding/json"
	"net/http"

	"github.com/ethanflaharty/Chirpy/internal/auth"
	"github.com/ethanflaharty/Chirpy/internal/database"
)

func (cfg *apiConfig) handlerAuthorization(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Password string `json:"password"`
		Email    string `json:"email"`
	}
	type response struct {
		User
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	tokenString, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Token is unauthorized", err)
		return
	}

	userID, err := auth.ValidateJWT(tokenString, cfg.secret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Unauthorized id", err)
		return
	}

	hashPass, err := auth.HashPassword(params.Password)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error hashing password", err)
		return
	}

	userUpdate, err := cfg.db.UpdateUser(r.Context(), database.UpdateUserParams{
		Email:          params.Email,
		HashedPassword: hashPass,
		ID:             userID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error updateing user", err)
		return
	}

	respondWithJSON(w, http.StatusOK, response{
		User: User{
			ID:        userUpdate.ID,
			CreatedAt: userUpdate.CreatedAt,
			UpdatedAt: userUpdate.UpdatedAt,
			Email:     userUpdate.Email,
		},
	})
}
