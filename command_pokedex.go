package main

import "fmt"

func commandPokedex(config *config, args ...string) error {
	if len(config.pokeDex) == 0 {
		fmt.Println("You have no pokemons in your PokeDex. Go catch some!")
		return nil
	}

	fmt.Println("You pokemon collection:")
	for key, _ := range config.pokeDex {
		fmt.Printf(" - %s\n", key)
	}

	return nil
}
