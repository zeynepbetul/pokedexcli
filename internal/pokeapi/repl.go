package pokeapi

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var url = "https://pokeapi.co/api/v2/location-area/"
var locations = LocationAreaResponse{}
var exploreResp = ExploreResponse{}

type cliCommand struct {
	name        string
	description string
	callback    func(string)
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
		"explore": {
			name:        "explore",
			description: "Explore the location area",
			callback:    commandExplore,
		},
	}
}

func commandMapb(areaToExplore string) {
	if locations.Previous != nil {
		locations.requestLocations(*locations.Previous)
		for index, location := range locations.Results {
			fmt.Println(index+1, ". location is:", location.Name)
		}
	} else {
		fmt.Println("You are on the first page")
	}
}

func commandMap(areaToExplore string) {
	if locations.Next == nil {
		locations.Next = &url // equalize address of the string to the locations.Next (it was in type *string)
	}
	locations.requestLocations(*locations.Next)
	for index, location := range locations.Results {
		fmt.Println(index+1, ". location is:", location.Name)
	}
}

func commandHelp(areaToExplore string) {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: ")
	fmt.Println(" ")

	myMap := getMap()
	for _, value := range myMap {
		fmt.Println(value.name, ": ", value.description)
	}
}

func commandExit(areaToExplore string) {
	os.Exit(0)
}

func commandExplore(areaToExplore string) {
	fmt.Println("Exploring", areaToExplore, "...")
	areaToExploreUrl := "https://pokeapi.co/api/v2/location-are" + areaToExplore // değişkede tutamıyorum neden?

	exploreResp.requestPokemons(areaToExploreUrl)
	fmt.Println("Found Pokemon:")
	for index, pokemon := range exploreResp.PokemonEncounters {
		fmt.Println(index+1, pokemon.Pokemon.Name) // neden pokemon.name değil?
	}
}

func StartRepl() {
	input := getUserInput()

	words := strings.Fields(input)

	myMap := getMap()

	value, found := myMap[words[0]]

	if found {
		if len(words) < 2 {
			words = append(words, "")
		}
		value.callback(words[1]) // diğer commandlere neden parametre veriyoruz bilmiyorum 6.explore da öyle istiyor buraya dön
	}
}

func getUserInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")

	scanner.Scan()
	input := scanner.Text()

	return input
}
