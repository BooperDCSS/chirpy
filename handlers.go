package main

import (
	"fmt"
	"log"
	"net/http"
)

func handlerCheckReady(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(http.StatusText(http.StatusOK)))
	if err != nil {
		log.Printf("%v\n", err)
	}
}

func (cfg *apiConfig) handlerRequestCount(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	message := fmt.Sprintf("Hits: %v\n", cfg.fileserverHits.Load())

	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(message))
	if err != nil {
		log.Printf("%v\n", err)
	}
}

func (cfg *apiConfig) handlerReset(w http.ResponseWriter, _ *http.Request) {
	cfg.fileserverHits.Store(0)

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Reset request received\n"))
	if err != nil {
		log.Printf("%v\n", err)
	}
}
