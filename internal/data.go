package internal

import "github.com/Muto1907/PokeDex/internal/pokecache"

type Location struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type Config struct {
	Next     string
	Previous *string
	Cache    *pokecache.Cache
}
