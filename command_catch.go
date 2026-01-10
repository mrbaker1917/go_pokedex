package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args []string) error {
	if len(args) != 1 {
		return errors.New("You must provide a pokemon name to catch.")
	}

	pokemonName := args[0]

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	poke, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	baseExperience := poke.BaseExperience

	roll := rand.Intn(baseExperience)

	if roll > 40 {
		fmt.Printf("%s escaped!\n", pokemonName)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemonName)
	fmt.Println("You may now inspect it with the inspect command.")

	cfg.pokedex[pokemonName] = poke

	return nil
}
