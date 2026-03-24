package main

import (
	"errors"
	"fmt"
)

func commandExplore(config *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	locationName := args[0]
	pokemons, err := config.pokeapiClient.GetPokemonsAtLocation(locationName)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", config.locationArea)
	fmt.Println("Found following Pokemons in the area: ")
	for _, result := range pokemons.Results {
		fmt.Printf(" - %s\n", result.Encounter.Name)
	}
	return nil
}
