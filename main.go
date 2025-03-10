package main


import (
	"log"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/ntwalibrian/urler/api"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", handler.Handler).Methods("GET")

	fmt.Printf("Starting server at port 8080 or on vercel \n")
	log.Fatal(http.ListenAndServe(":8080", r))
}