package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *config) error {
	fmt.Println("Exiting the Pokedex. Goodbye!")
	os.Exit(0)
	return nil
}
