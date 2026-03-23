package main

import (
	"time"

	"github.com/vmarin93/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	config := &config{
		pokeapiClient: pokeClient,
	}
	pokeDex(config)
}
