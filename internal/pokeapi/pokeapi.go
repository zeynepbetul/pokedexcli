package pokeapi

import (
	"encoding/json"
	"fmt"
	"github.com/zeynepbetul/pokedexcli/internal/pokecache"
	"io"
	"log"
	"net/http"
)

func (locationsStruct *LocationAreaResponse) requestLocations(endpointUrl string) { // receiver should be type 'pointer receiver' to be able to use original struct. Otherwise, it creates a copy of it and make changes on the copy.
	// function to be part of the struct's methods OR being a standalone function that can operate on any instance of LocationAreaResponse passed as a pointer. -- OPEN
	var body []byte
	pokecache.FirstCache.Mu.Lock()
	entry, ok := pokecache.FirstCache.Get(endpointUrl) // try to get response body from the cache
	pokecache.FirstCache.Mu.Unlock()

	if ok {
		body = entry
	} else {
		response, err := http.Get(endpointUrl)
		if err != nil {
			log.Fatal(err)
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				fmt.Println("Error close body", err)
			}
		}(response.Body)
		body, err = io.ReadAll(response.Body) // body --> slice of bytes --> []byte --> [78 111 116 32 70 111 117 110 100]
		if err != nil {
			fmt.Println("Error reading response body", err)
		}
		pokecache.FirstCache.Mu.Lock()
		pokecache.FirstCache.Add(endpointUrl, body) // cache all the response body which have: count, next, prev, results in it
		pokecache.FirstCache.Mu.Unlock()
	}
	if err := json.Unmarshal(body, &locationsStruct); err != nil {
		fmt.Println(err)
	}
}

func (exploreStruct *ExploreResponse) requestPokemons(endpointUrl string) {
	response, err := http.Get(endpointUrl)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body")
	}

	if err := json.Unmarshal(body, &exploreStruct); err != nil {
		fmt.Println(err)
	}
}
