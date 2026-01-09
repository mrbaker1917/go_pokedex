package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args []string) error {
	if len(args) == 0 {
		return errors.New("You must provide a pokemon name to catch.")
	}

	pokemonName := args[0]

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	poke, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}
	roll := rand.Intn(100)

	baseExperience := poke.BaseExperience
	catchChance := 80 - (baseExperience / 5)

	if catchChance < 5 {
		catchChance = 5
	}

	if catchChance > 95 {
		catchChance = 95
	}

	if roll < catchChance {
		fmt.Printf("%s was caught!\n", pokemonName)
		cfg.pokedex[pokemonName] = poke
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}

	return nil
}
