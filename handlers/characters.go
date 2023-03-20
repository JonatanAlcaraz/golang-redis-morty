package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang-redis-morty/connection"
	"golang-redis-morty/service"

	"github.com/go-redis/redis"
)

func GetCharactersHandler(w http.ResponseWriter, r *http.Request) {
	client := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})
	defer client.Close()

	key := "rick-and-morty-characters"

	characters, err := connection.GetCharacters(client, key)
	if err != nil {
		fmt.Println("No se encontro la siguiente key en redis:", key)
		c := make(chan []*connection.Character)
		go service.GetRickAndMortyCharacters("https://rickandmortyapi.com/api/character", c)
		characters = <-c

		connection.SaveCharacters(client, key, characters)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(characters)
}
