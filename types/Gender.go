package pokego

// Gender ...
type Gender struct {
	ID                    int    `json:"id"`
	Name                  string `json:"name"`
	PokemonSpeciesDetails []struct {
		PokemonSpecies struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon_species"`
		Rate int `json:"rate"`
	} `json:"pokemon_species_details"`
	RequiredForEvolution []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"required_for_evolution"`
}
