package main

import "fmt"

func commandExplore(config *config) error {
	pokemons, err := config.pokeapiClient.GetPokemonsAtLocation(config.locationArea)
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
