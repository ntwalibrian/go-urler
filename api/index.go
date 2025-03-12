package api

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

type RequestBody struct {
	URL string `json:"url"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1> Hello vercel by diouf </h1>")
}

func Shorten(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var requestBody RequestBody
	err := json.NewDecoder(r.Body).Decode(&requestBody)

	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Validate if URL is present
	if requestBody.URL == "" {
		http.Error(w, "Missing URL in request body", http.StatusBadRequest)
		return
	}

	short := RandKey(8)

	FileWrite(short, requestBody.URL)

	// Send response
	json.NewEncoder(w).Encode(map[string]string{
		"short_url": "http://localhost:8080/" + short,
	})


}

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortKey := vars["shortKey"]

	urlMappings := FileRead()

	// Check if the short key exists
	if originalURL, exists := urlMappings[shortKey]; exists {
		http.Redirect(w, r, originalURL, http.StatusFound) // 302 Redirect
	} else {
		http.Error(w, "URL not found", http.StatusNotFound)
	}
}
