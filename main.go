package main

import (
	"time"

	pokeapi "github.com/HiroAcocoro/go-pokedex/internal/api"
)

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
	}

	startPokeRepl(&cfg)
}
