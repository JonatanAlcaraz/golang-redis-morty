package conection

import (
	"encoding/json"

	"github.com/go-redis/redis"
)

type Character struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Image string `json:"image"`
}

func GetCharacters(client *redis.Client, key string) ([]*Character, error) {
    val, err := client.Get(key).Result()
    if err != nil {
        return nil, err
    }

    var characters []*Character
    err = json.Unmarshal([]byte(val), &characters)
    if err != nil {
        return nil, err
    }

    return characters, nil
}

func SaveCharacters(client *redis.Client, key string, characters []*Character) error {
    data, err := json.Marshal(characters)
    if err != nil {
        return err
    }

    err = client.Set(key, data, 0).Err()
    if err != nil {
        return err
    }

    return nil
}