package service

import (
	"encoding/json"
	"golang-redis-morty/connection"
	"log"
	"net/http"
)


func GetRickAndMortyCharacters(url string, c chan<- []*connection.Character) {
    resp, err := http.Get(url)
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()

    var response struct {
        Results []*connection.Character `json:"results"`
    }
    err = json.NewDecoder(resp.Body).Decode(&response)
    if err != nil {
        log.Fatal(err)
    }

    c <- response.Results
}


