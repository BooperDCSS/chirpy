package main

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
)

func main() {
	const port = "8080"
	rootPath := "."

	apiCfg := &apiConfig{
		fileserverHits: atomic.Int32{},
	}

	mux := http.NewServeMux()

	handlerApp := http.StripPrefix("/app", http.FileServer(http.Dir(rootPath)))
	mux.Handle("/app/", apiCfg.middlewareMetricsInc(handlerApp))

	mux.HandleFunc("GET /api/healthz", handlerCheckReady)
	mux.HandleFunc("GET /api/metrics", apiCfg.handlerRequestCount)
	mux.HandleFunc("POST /api/reset", apiCfg.handlerReset)

	myServer := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	fmt.Printf("Serving on port %q\n", port)

	err := myServer.ListenAndServe()
	if err != nil {
		log.Fatalf("Error on myServer: %v", err)
	}

}
