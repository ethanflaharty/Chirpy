package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func (cfg *apiConfig) handlerValidate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		log.Printf("Error decoding parameters: %s", err)
		w.WriteHeader(500)
		return
	}

	if len(params.Body) > 140 {
		resp := errResponse{
			Error: "Chirp is too long",
		}

		dat, err := json.Marshal(resp)
		if err != nil {
			w.WriteHeader(500)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		w.Write(dat)
		return
	}

	clean_body := censorWords(params.Body)

	resp := validResponse{
		CleanedBody: clean_body,
	}

	dat, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(dat)
}

type validResponse struct {
	CleanedBody string `json:"cleaned_body"`
}

type errResponse struct {
	Error string `json:"error"`
}

func censorWords(s string) string {
	splitUp := strings.Split(s, " ")
	for i, word := range splitUp {
		word = strings.ToLower(word)
		if word == "kerfuffle" || word == "sharbert" || word == "fornax" {
			splitUp[i] = "****"
		}
	}
	cleaned := strings.Join(splitUp, " ")
	return cleaned
}
