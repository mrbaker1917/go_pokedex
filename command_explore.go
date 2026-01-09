package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args []string) error {
	if len(args) == 0 {
		return errors.New("You must provide a location area name")
	}

	areaName := args[0]

	fmt.Printf("Exploring %s...\n", areaName)

	loc, err := cfg.pokeapiClient.GetLocationArea(areaName)
	if err != nil {
		return err
	}
	fmt.Println("Found Pokemon:")
	for _, enc := range loc.PokemonEncounters {
		fmt.Printf(" - %s\n", enc.Pokemon.Name)
	}
	return nil
}
