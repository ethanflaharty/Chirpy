package main

import (
	"log"
	"net/http"
)

func (cfg *apiConfig) handlerReset(w http.ResponseWriter, r *http.Request) {
	if cfg.platform != "dev" {
		w.WriteHeader(403)
		return
	}
	err := cfg.query.DeleteUsers(r.Context())
	if err != nil {
		log.Printf("Error deleting users: %s", err)
		w.WriteHeader(500)
		return
	}
	cfg.fileserverHits.Store(0)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hits reset to 0"))
}
