package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"html/template"
	"github.com/gorilla/mux"
	"github.com/ntwalibrian/urler/api"
)

func resultPageHandler(w http.ResponseWriter, r *http.Request) {
	// Get the query parameters (original_url and short_url)
	originalURL := r.URL.Query().Get("original_url")
	shortURL := r.URL.Query().Get("short_url")

	// If the URLs are not provided, send an error message
	if originalURL == "" || shortURL == "" {
		http.Error(w, "Missing URL parameters", http.StatusBadRequest)
		return
	}

	// Define the data to pass to the template
	data := struct {
		OriginalURL string
		ShortURL    string
	}{
		OriginalURL: originalURL,
		ShortURL:    shortURL,
	}

	// Parse the result.html template from the static folder
	tmpl, err := template.ParseFiles("./static/responce.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	tmpl.Execute(w, data)
}


func main() {
	r := mux.NewRouter()
	
	r.HandleFunc("/shorten", api.Shorten).Methods("POST")
	r.HandleFunc("/webshorten", api.WebShorten).Methods("POST")
	r.HandleFunc("/result", resultPageHandler).Methods("GET")
	r.HandleFunc("/{shortKey}", api.RedirectHandler).Methods("GET")
	
	fileServer := http.FileServer(http.Dir("./static"))
    r.PathPrefix("/").Handler(fileServer)

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