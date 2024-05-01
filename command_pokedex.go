package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {

	fmt.Println("Your Pokedex:")
	for p := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", p)
	}
	return nil
}
