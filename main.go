package main

import (
	"math/rand"
	"time"

	"github.com/mrbaker1917/go_pokedex/internal/pokeapi"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	pokeClient := pokeapi.NewClient(
		5*time.Second,
		5*time.Second,
	)
	cfg := &config{
		pokeapiClient: pokeClient,
		pokedex:       make(map[string]pokeapi.Pokemon),
	}

	startRepl(cfg)
}
