package service

import (
	"encoding/json"
	"golang-redis-morty/conection"
	"log"
	"net/http"
)


func GetRickAndMortyCharacters(url string, c chan<- []*conection.Character) {
    resp, err := http.Get(url)
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()

    var response struct {
        Results []*conection.Character `json:"results"`
    }
    err = json.NewDecoder(resp.Body).Decode(&response)
    if err != nil {
        log.Fatal(err)
    }

    c <- response.Results
}


