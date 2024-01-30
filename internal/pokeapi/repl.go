package pokeapi

import (
	"bufio"
	"fmt"
	"os"
)

var url = "https://pokeapi.co/api/v2/location-area/"
var locations = LocationAreaResponse{}

type cliCommand struct {
	name        string
	description string
	callback    func()
}

func getMap() map[string]cliCommand {

	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Show next map",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Show previous map",
			callback:    commandMapb,
		},
	}
}

func commandMapb() {
	if locations.Previous != nil {
		locations.requestLocations(*locations.Previous)
		for index, location := range locations.Results {
			fmt.Println(index+1, ". location is:", location.Name)
		}
	} else {
		fmt.Println("You are on the first page")
	}
}

func commandMap() {
	if locations.Next == nil {
		locations.Next = &url // equalize address of the string to the locations.Next (it was in type *string)
	}
	locations.requestLocations(*locations.Next)
	for index, location := range locations.Results {
		fmt.Println(index+1, ". location is:", location.Name)
	}
}

func commandHelp() {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: ")
	fmt.Println(" ")

	myMap := getMap()
	for _, value := range myMap {
		fmt.Println(value.name, ": ", value.description)
	}

}

func commandExit() {
	os.Exit(0)
}

func StartRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")
	scanner.Scan()
	text := scanner.Text()

	myMap := getMap()

	value, found := myMap[text]

	if found {
		value.callback()
	}
}
