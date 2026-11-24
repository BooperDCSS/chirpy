package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	const port = "8080"

	mux := http.NewServeMux()
	mux.Handle("/app/", http.StripPrefix("/app",
		http.FileServer(http.Dir("."))))

	mux.HandleFunc("/healthz", handlerCheckReady)
	
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

func handlerCheckReady(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(http.StatusText(http.StatusOK)))
	if err != nil {
		log.Printf("%v\n", err)
		}
}
