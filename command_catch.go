package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(config *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	pokemonName := args[0]
	if _, ok := config.pokeDex[pokemonName]; ok {
		fmt.Printf("%s is already in your PokeDex\n", pokemonName)
		return nil
	}

	pokemon, err := config.pokeapiClient.CatchPokemon(pokemonName)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	chance := rand.Intn(pokemon.BaseExperience)
	benchmark := pokemon.BaseExperience / 2

	if chance >= benchmark {
		config.pokeDex[pokemon.Name] = pokemon
		fmt.Printf("%s was cought!\n", pokemon.Name)
		fmt.Println("You may now inspect it with the inspect command.")
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}
	return nil
}
