package main

import (
	"fmt"
	"golang-redis-morty/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
    
	r.HandleFunc("/characters", handlers.GetCharactersHandler).Methods("GET")
	
	fmt.Println("Listening on port 3000")
	http.ListenAndServe(":3000", r)
}
