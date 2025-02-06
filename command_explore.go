package main

import (
	"fmt"
)


func commandExplore(c *config, args []string) error{
	if len(args) < 2{
		return fmt.Errorf("err no area listed")
	}

	areaName := args[1]

	pokemon, err := c.pokeapiClient.ExploreLocation(areaName)
	if err != nil {
		return fmt.Errorf("%w",err)
	}

	for _, v := range pokemon{
		fmt.Println("-",v.Name)
	}

	return nil
}
