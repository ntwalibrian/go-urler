package main

import (
	"log"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	fmt.Printf("Starting server at port 8080 or on vercel \n")
	log.Fatal(http.ListenAndServe(":8080", r))
}