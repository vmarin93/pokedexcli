package main

import (
	"fmt"
	"strings"

	"github.com/vmarin93/pokedexcli/internal/pokeapi"
)

func pokeDex(config *config) {
	commands := getCommands()
	prompt := UserPrompt{
		inputHistory: []string{""},
	}
	for {
		input, err := prompt.getUserInput()
		if err != nil {
			fmt.Println(err)
		}
		if len(input) == 0 {
			continue
		}
		cmdName := input[0]
		args := []string{}
		if len(input) > 1 {
			args = append(args, input[1])
		}
		cmd, ok := commands[cmdName]
		if ok {
			err := cmd.callback(config, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

type config struct {
	pokeapiClient   pokeapi.Client
	nextLocationUrl *string
	prevLocationUrl *string
	pokeDex         map[string]pokeapi.Pokemon
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in the Pokemon world.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of the previous 20 location areas in the Pokemon world.",
			callback:    commandMapBack,
		},
		"explore": {
			name:        "explore <area-name>",
			description: "Displays the names of all Pokemon that can be found in a specific area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Attempts to catch a pokemon and add it to your PokeDex",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "Inspect a Pokemon in your PokeDex to see more information about it",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List all the pokemons that you have caught inside your PokeDex",
			callback:    commandPokedex,
		},
	}
}
