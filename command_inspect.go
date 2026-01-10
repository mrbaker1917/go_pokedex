package main

import (
	"fmt"
)

func commandInspect(cfg *config, args []string) error {
	pokemonName := args[0]
	for _, p := range cfg.pokedex {
		if pokemonName == p.Name {
			fmt.Printf("Name: %s\nHeight: %d\nWeight: %d\n", p.Name, p.Height, p.Weight)
			fmt.Println("Stats:")
			for _, stat := range p.Stats {
				fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
			}
			fmt.Println("Types:")
			for _, t := range p.Types {
				fmt.Printf("  -%s\n", t.Type.Name)
			}
			fmt.Println()
			return nil
		}
	}
	fmt.Printf("You have not caught that pokemon yet!\n")
	return nil
}
