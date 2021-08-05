package main

import "fmt"

/*
 * Interfaces are named collections of method signatures.
 *
 * An interface is two things: it is a set of methods, but it is also a type.
 *
 * References by:
 * https://gobyexample.com/interfaces
 * https://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go
 */
type (
	Pokemon interface {
		Type() string
	}

	Charizard struct {
		typ       string
		evolution string
	}

	Venusaur struct {
		typ     string
		evo     string
		pokedex string
	}

	Blastoise struct {
	}
)

func (c Charizard) Type() string {
	c.typ = "Fire"
	c.evolution = "3ª"
	msg := "Charizard is the " + c.typ + " type and " + c.evolution + " evolution."
	return msg
}

func (v Venusaur) Type() string {
	v.typ = "Grass"
	v.evo = "3ª"
	v.pokedex = "3"
	msg := "Venusaur is the " + v.typ + " type, " + v.evo + " evolution, and number " + v.pokedex + " in pokedex."
	return msg
}

func (b Blastoise) Type() string {
	return "Blastoise is the water type, 3ª evolution, and number 9 in pokedex."
}

func main() {
	pokemons := []Pokemon{Charizard{}, Venusaur{}, Blastoise{}}
	for _, pokemon := range pokemons {
		fmt.Println(pokemon.Type())
	}
}
