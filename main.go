package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
}

func cleanInput(text string) []string {
	text = strings.TrimSpace(text)
	fmt.Println(text)

	x := strings.Fields(text)
	fmt.Println(x)
	return x
}
