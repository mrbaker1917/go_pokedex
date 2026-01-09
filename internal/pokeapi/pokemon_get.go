package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(name string) (PokemonInfo, error) {
	url := baseURL + "/pokemon/" + name

	if dat, ok := c.cache.Get(url); ok {
		pokemonResp := PokemonInfo{}
		if err := json.Unmarshal(dat, &pokemonResp); err != nil {
			return PokemonInfo{}, err
		}
		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonInfo{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonInfo{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		return PokemonInfo{}, fmt.Errorf("Pokemon %s not found", name)
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonInfo{}, err
	}

	c.cache.Add(url, dat)

	pokemonResp := PokemonInfo{}
	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return PokemonInfo{}, err
	}

	return pokemonResp, nil
}
