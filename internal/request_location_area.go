package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (cl *Client) Request_location_area(area string) (Location_area, error) {
	url := baseURL + "/location-area/" + area
	body, ok := cl.cache.Get(url)
	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return Location_area{}, err
		}
		res, err := cl.httpClient.Do(req)
		if err != nil {
			return Location_area{}, err
		}
		if res.StatusCode == 404 {
			return Location_area{}, fmt.Errorf("please enter a valid location area, check map for more locations")
		}
		defer res.Body.Close()
		body, err = io.ReadAll(res.Body)
		if err != nil {
			return Location_area{}, err
		}
		cl.cache.Add(url, body)
	}
	location_area := Location_area{}
	err := json.Unmarshal(body, &location_area)
	if err != nil {
		return Location_area{}, err
	}
	return location_area, nil
}
