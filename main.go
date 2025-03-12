package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/ntwalibrian/urler/api"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", api.Handler).Methods("GET")
	r.HandleFunc("/shorten", api.Shorten).Methods("POST")
	r.HandleFunc("/{shortKey}", api.RedirectHandler).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Starting server at port 8080 or on Render \n")

	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}