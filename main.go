package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	pokeapi "github.com/HiroAcocoro/go-pokedex/internal/api"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
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
			description: "List location areas",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "List previous location ares",
			callback:    callbackMapb,
		},
	}
}
func printPrompt() {
	fmt.Print("pokedex >")
}

func sanitizeInput(input string) string {
	output := strings.TrimSpace(input)
	output = strings.ToLower(output)
	return output
}

func commandSelector(cmd string, cfg *config) {
	commands := getCommands()
	if command, ok := commands[cmd]; ok {
		err := command.callback(cfg)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("Uknown command. Type 'help' for a list of avaialable commands.")
	}
}

func commandHelp(cfg *config) error {
	commands := getCommands()
	fmt.Println("Welcome to the Pokedex!\nUsage:\n\n")
	for _, cmd := range commands {
		fmt.Printf("%s: %s", cmd.name, cmd.description)
		fmt.Println("")
	}
	return nil
}

func commandExit(cfg *config) error {
	fmt.Println("Exiting the Pokedex. Goodbye!")
	os.Exit(0)
	return nil
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(),
	}
	reader := bufio.NewScanner(os.Stdin)

	for {
		printPrompt()
		if !reader.Scan() {
			break
		}

		cmd := sanitizeInput(reader.Text())
		commandSelector(cmd, &cfg)
	}

}
