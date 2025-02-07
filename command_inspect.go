package main

import (
	"fmt"

	"github.com/Ekaloi/pokedex/poketypes"
)

func commandInspect(c *config, args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("no pokemon specified")
	}

	pokemon, err := c.pokeapiClient.Cache.Inspect(args[1])
	if err != nil {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	printPokemon(pokemon)
	return nil
}


func printPokemon(pokemon poketypes.PokemonActual) error {
	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats:")
	for _, v := range pokemon.Stats{
		fmt.Printf("-%s: %d\n",v.Stat.Name, v.BaseStat)
	}
	fmt.Println("Types:")
	for _, v := range pokemon.Types{
		fmt.Printf("- %s\n",v.Type.Name)
	}
	return nil
}

