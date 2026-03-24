package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/vmarin93/pokedexcli/internal/pokeapi"
)

func pokeDex(config *config) {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		prompt := cleanInput(scanner.Text())
		if len(prompt) == 0 {
			continue
		}
		cmdName := prompt[0]
		args := []string{}
		if len(prompt) > 1 {
			args = prompt[1:]
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
			name:        "explore",
			description: "Displays the names of all Pokemon that can be found at location <location name>",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Throw a pokeball at a Pokemon in the attempts to catch it. Any succesfull attempt will add the Pokemon to your PokeDex",
			callback:    commandCatch,
		},
	}
}
