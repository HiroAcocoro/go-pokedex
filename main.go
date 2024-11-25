package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
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

func commandSelector(cmd string) {
	commands := getCommands()
	if command, ok := commands[cmd]; ok {
		command.callback()
	} else {
		fmt.Println("Uknown command. Type 'help' for a list of avaialable commands.")
	}
}

func commandHelp() error {
	commands := getCommands()
	fmt.Println("Welcome to the Pokedex!\nUsage:\n\n")
	for _, cmd := range commands {
		fmt.Printf("%s: %s", cmd.name, cmd.description)
		fmt.Println("")
	}
	return nil
}

func commandExit() error {
	fmt.Println("Exiting the Pokedex. Goodbye!")
	os.Exit(0)
	return nil
}

func main() {
	reader := bufio.NewScanner(os.Stdin)

	for {
		printPrompt()
		if !reader.Scan() {
			break
		}

		cmd := sanitizeInput(reader.Text())
		commandSelector(cmd)
	}

}
