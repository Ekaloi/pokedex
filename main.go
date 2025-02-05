package main

import "github.com/Ekaloi/pokedex/internal/pokeapi"


func main() {
	pokeClient := pokeapi.NewClient()
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}





