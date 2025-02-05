package main


import (
	"fmt"
	
)

func commandHelp(c *config) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	fmt.Println()

	for _, v := range getCommands() {
		fmt.Printf("%s: %s\n", v.name, v.description)
	}
	return nil
}