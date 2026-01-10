package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args []string) error {
	if len(cfg.pokedex) < 1 {
		fmt.Printf("You have not caught any pokemon yet!\n")
		return nil
	}
	fmt.Println("Your Pokedex:")

	for _, p := range cfg.pokedex {
		fmt.Printf("  -%s\n", p.Name)
	}
	return nil
}
