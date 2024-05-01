package main

import "fmt"

func commandHelp(cfg *config, args ...string) error {
	fmt.Println("Welcome to the Pokedex! \n Usage:\n help:Displays a help message\n exit: Exit the Pokedex")
	return nil
}
