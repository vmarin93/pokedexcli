package main

import (
	"errors"
	"fmt"
)

func commandInspect(config *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("please provide a pokemon name")
	}

	pokemonName := args[0]
	if _, ok := config.pokeDex[pokemonName]; !ok {
		err := fmt.Errorf("%s is not in your PokeDex yet. You have to catch it first", pokemonName)
		return err
	}
	pokemon := config.pokeDex[pokemonName]

	fmt.Printf("Name: %s\nBase Experience: %d\nHeight: %d\nWeight: %d\n", pokemon.Name, pokemon.BaseExperience, pokemon.Height, pokemon.Weight)
	fmt.Println("Abilities:")
	for _, ability := range pokemon.Abilities {
		fmt.Printf(" - %s\n", ability.Ability.Name)
	}
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	return nil
}
