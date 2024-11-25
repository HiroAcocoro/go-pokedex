package main

import "fmt"

func commandHelp(cfg *config) error {
	commands := getCommands()
	fmt.Println("Welcome to the Pokedex!\nUsage:\n\n")
	for _, cmd := range commands {
		fmt.Printf("%s: %s", cmd.name, cmd.description)
		fmt.Println("")
	}
	return nil
}
