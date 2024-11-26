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

func sanitizeInput(input string) []string {
	output := strings.ToLower(input)
	words := strings.Fields(output)
	return words
}

func commandSelector(cmd string, cfg *config, args ...string) {
	commands := getCommands()

	command, ok := commands[cmd]
	if !ok {
		fmt.Println("Uknown command. Type 'help' for a list of avaialable commands.")
		return
	}

	err := command.callback(cfg, args...)
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

		cleanCmd := sanitizeInput(reader.Text())

		if len(cleanCmd) == 0 {
			continue
		}

		cmd := cleanCmd[0]
		args := []string{}
		if len(cleanCmd) > 1 {
			args = cleanCmd[1:]
		}
		commandSelector(cmd, cfg, args...)
	}
}
