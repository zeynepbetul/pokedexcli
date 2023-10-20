package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type LocationAreaResponse struct {
	Count    int      `json:"count"`
	Next     *string  `json:"next"`
	Previous *string  `json:"previous"`
	Results  []Result `json:"results"`
}

type Result struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func (locations *LocationAreaResponse) requestLocations(endpointUrl string) { // receiver should be type 'pointer receiver' to be able to use original struct. Otherwise, it creates a copy of it and make changes on the copy.
	// function to be part of the struct's methods OR being a standalone function that can operate on any instance of LocationAreaResponse passed as a pointer. -- OPEN
	response, err := http.Get(endpointUrl)

	if err != nil {
		// fmt.Println("Error making get request", err)
		// return
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body) // body --> slice of bytes -> []byte -> [78 111 116 32 70 111 117 110 100]
	if err != nil {
		fmt.Println("Error reading response body", err)
	}

	if err := json.Unmarshal(body, &locations); err != nil { // in one line. Use same variable name multiple times.
		fmt.Println(err)
	}
}
