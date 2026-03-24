package main

import (
	"fmt"
)

func commandMap(config *config, args ...string) error {
	locations, err := config.pokeapiClient.GetLocationsList(config.nextLocationUrl)
	if err != nil {
		return err
	}

	config.nextLocationUrl = locations.Next
	config.prevLocationUrl = locations.Previous

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapBack(config *config, args ...string) error {
	locations, err := config.pokeapiClient.GetLocationsList(config.prevLocationUrl)
	if err != nil {
		return err
	}

	config.nextLocationUrl = locations.Next
	config.prevLocationUrl = locations.Previous

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	return nil
}
