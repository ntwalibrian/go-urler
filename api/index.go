package api

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

type RequestBody struct {
	URL string `json:"url"`
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
		"short_url": "https://go-urler.onrender.com/" + short,
	})

}

func WebShorten(w http.ResponseWriter, r *http.Request) {

	// Parse the form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	// Extract URL from form data
	url := r.FormValue("url")
	if url == "" {
		http.Error(w, "Missing URL in form data", http.StatusBadRequest)
		return
	}

	short := RandKey(8)

	FileWrite(short, url)

	http.Redirect(w, r, "/result?original_url=" + url + "&short_url=https://go-urler.onrender.com/" + short, http.StatusFound)

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
