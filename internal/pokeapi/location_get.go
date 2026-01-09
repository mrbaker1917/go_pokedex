package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationArea(name string) (LocationArea, error) {
	url := baseURL + "/location-area/" + name

	if dat, ok := c.cache.Get(url); ok {
		locationResp := LocationArea{}
		if err := json.Unmarshal(dat, &locationResp); err != nil {
			return LocationArea{}, err
		}
		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		return LocationArea{}, fmt.Errorf("Location area %s not found", name)
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}

	c.cache.Add(url, dat)

	locationResp := LocationArea{}
	err = json.Unmarshal(dat, &locationResp)
	if err != nil {
		return LocationArea{}, err
	}

	return locationResp, nil
}
