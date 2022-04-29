package main

import (
	"fmt"
)

/*
 * Interfaces examples
 *
 * References by:
 * https://en.wikipedia.org/wiki/Duck_typing
 */
type (
	Animal interface {
		swim() string
		fly() string
	}

	Duck struct {}

	Whale struct {}
)

func (d Duck) swim() string {
	return "Duck swimming"
}

func (d Duck) fly() string {
	return "Duck flying"
}

func (w Whale) swim() string {
	return "Whale swimming"
}

func (w Whale) fly() string {
	return "Whale does not fly"
}

func main() {
	animals := []Animal{Duck{}, Whale{}}
	for _, animal := range animals {
		fmt.Println(animal.swim())
		fmt.Println(animal.fly())
	}
}