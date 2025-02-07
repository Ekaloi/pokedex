package main

import "fmt"

func commandPokedex(c *config, args []string) error {
	fmt.Println("Your Pokedex:")

	pokemon, err := c.pokeapiClient.Cache.InspectPokedex()
	if err != nil {
		return fmt .Errorf("error issue retreiving pokmemon from cache")
	}
	for _,v := range pokemon{
		fmt.Println(" - ", v.Name)
	}
	return nil
}