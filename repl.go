package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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

	command, ok := commands[cmd]
	if !ok {
		fmt.Println("Uknown command. Type 'help' for a list of avaialable commands.")
		return
	}

	err := command.callback(cfg)
	if err != nil {
		fmt.Println(err)
	}
}

func startPokeRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		printPrompt()
		if !reader.Scan() {
			break
		}

		cmd := sanitizeInput(reader.Text())
		commandSelector(cmd, cfg)
	}
}
