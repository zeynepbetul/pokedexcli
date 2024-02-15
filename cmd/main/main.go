package main

import "github.com/zeynepbetul/pokedexcli/internal/pokeapi"

func main() { // go build && ./pokedexcli
	for {
		pokeapi.StartRepl()
	}
}
