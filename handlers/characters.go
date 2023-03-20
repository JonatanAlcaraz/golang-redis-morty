package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang-redis-morty/conection"
	"golang-redis-morty/service"

	"github.com/go-redis/redis"
)

func GetCharactersHandler(w http.ResponseWriter, r *http.Request) {
	client := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})
	defer client.Close()

	key := "rick-and-morty-characters"

	characters, err := conection.GetCharacters(client, key)
	if err != nil {
		fmt.Println("No se encontro la siguiente key en redis:", key)
		c := make(chan []*conection.Character)
		go service.GetRickAndMortyCharacters("https://rickandmortyapi.com/api/character", c)
		characters = <-c

		conection.SaveCharacters(client, key, characters)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(characters)
}
