package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

type locationarea struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMap(cfg *Config) error {

	offstr := strconv.Itoa(cfg.CommandCalls["map"] * 20)
	locURL := "https://pokeapi.co/api/v2/location-area/?offset=" + offstr

	res, err := http.Get(locURL)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	locations := locationarea{}
	err = json.Unmarshal(body, &locations)
	if err != nil {
		fmt.Println(err)
	}
	for _, l := range locations.Results {
		fmt.Println(l.Name)
	}

	return nil
}

func commandMapB(cfg *Config) error {
	if cfg.CommandCalls["map"] > 0 {
		cfg.CommandCalls["map"] -= 2
		commandMap(cfg)
	} else {
		commandMap(cfg)
	}
	return nil
}
