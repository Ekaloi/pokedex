package main

import (
	"fmt"
	"math/rand"
	"time"
)

func commandCatch(c *config, args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("no pokemon specified")
	}
	fmt.Println("Throwing a Pokeball at " + args[1] + "...")

	pokemon, err := c.pokeapiClient.GetPokemon(args[1])
	if err != nil {
		return err
	}
	rand.Seed(time.Now().UnixNano())
	val := catch(pokemon.BaseExperience)
	if val {
		fmt.Printf("%s was caught!\n", args[1])
		return nil
	} else {
		fmt.Printf("%s escaped!\n", args[1])
		return nil
	}
}

func catch(x int) bool {
	// rand.Intn(10) returns a number in [0,9].
	// For a 1 in 10 chance, we can return true if the result is 0.
	return rand.Intn(10)*x < 1000
}
