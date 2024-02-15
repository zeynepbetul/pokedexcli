package pokeapi

type ExploreResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	GameIndex int    `json:"game_index"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

/*
In Go, those backticks (`) are used to denote struct tags.
Struct tags are metadata attached to struct fields that provide additional information about the field.
They are often used to specify how a struct field should be serialized or deserialized when working with formats like JSON, XML, or database tables.

In your example, the struct tags are used to specify how the fields of the ExploreResponse struct should be mapped to JSON keys when
marshaling (encoding) and unmarshaling (decoding) JSON data.
For example:

json:"id" indicates that the ID field should be mapped to the JSON key "id".
json:"name" indicates that the Name field should be mapped to the JSON key "name".
json:"location" and json:"pokemon_encounters" are used for nested structs, indicating how those nested structs should be represented in JSON.
*/
