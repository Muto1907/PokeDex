package internal

import (
	"encoding/json"
	"fmt"
	"io"
)

func (cl *Client) Request_locations(pageURL *string) (Location, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}
	body, ok := cl.cache.Get(url)
	if !ok {
		res, err := cl.httpClient.Get(url)
		if err != nil {
			return Location{}, err
		}
		body, err = io.ReadAll(res.Body)
		defer res.Body.Close()
		if res.StatusCode > 299 {
			return Location{}, fmt.Errorf("response failed with status code: %d and \n body: %s", res.StatusCode, body)
		}
		if err != nil {
			return Location{}, err
		}
		cl.cache.Add(url, body)
	}

	location := Location{}
	err := json.Unmarshal(body, &location)
	if err != nil {
		return Location{}, err
	}
	return location, nil
}
