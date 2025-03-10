package main

import (
	"log"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello vercel by diouf </h1>")
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", helloHandler).Methods("GET")

	fmt.Printf("Starting server at port 8080 or on vercel \n")
	log.Fatal(http.ListenAndServe(":8080", r))
}