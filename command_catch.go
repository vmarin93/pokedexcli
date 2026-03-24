package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(config *config) error {
	if _, ok := config.pokeDex[config.pokemonName]; ok {
		fmt.Printf("%s is already in your PokeDex\n", config.pokemonName)
		return nil
	}

	pokemon, err := config.pokeapiClient.CatchPokemon(config.pokemonName)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	chance := rand.Intn(pokemon.BaseExperience)
	benchmark := pokemon.BaseExperience / 2

	if chance >= benchmark {
		config.pokeDex[pokemon.Name] = pokemon
		fmt.Printf("%s was cought!\n", config.pokemonName)
	} else {
		fmt.Printf("%s escaped!\n", config.pokemonName)
	}
	return nil
}
