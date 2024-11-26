package main

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
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
		"explore": {
			name:        "explore {location_area}",
			description: "List previous location ares",
			callback:    callbackExplore,
		},
	}
}
