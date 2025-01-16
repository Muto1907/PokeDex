package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (cl *Client) Request_pokemon(pokemon_name string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemon_name
	body, ok := cl.cache.Get(url)
	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return Pokemon{}, err
		}
		res, err := cl.httpClient.Do(req)
		if err != nil {
			return Pokemon{}, err
		}
		defer res.Body.Close()
		body, err = io.ReadAll(res.Body)
		if err != nil {
			return Pokemon{}, err
		}
		if res.StatusCode == 404 {
			return Pokemon{}, fmt.Errorf("pokemon %s does not exist", pokemon_name)
		}
		cl.cache.Add(url, body)
	}
	pokemon := Pokemon{}
	err := json.Unmarshal(body, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}
	return pokemon, nil
}
