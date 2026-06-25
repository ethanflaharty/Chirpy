package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerUsers(w http.ResponseWriter, r *http.Request) {
	params := struct {
		Email string `json:"email"`
	}{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		log.Printf("Error decoding users: %s", err)
		w.WriteHeader(500)
		return
	}

	dbUsers, err := cfg.query.CreateUser(r.Context(), params.Email)
	if err != nil {
		log.Printf("Error creating users: %s", err)
		w.WriteHeader(500)
		return
	}

	users := User{
		ID:        dbUsers.ID,
		CreatedAt: dbUsers.CreatedAt,
		UpdatedAt: dbUsers.UpdatedAt,
		Email:     dbUsers.Email,
	}

	dat, err := json.Marshal(users)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(201)
	w.Write(dat)
}

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
}
