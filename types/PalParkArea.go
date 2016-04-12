package pokego

// PalParkArea
type PalParkArea struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		BaseScore      int `json:"base_score"`
		PokemonSpecies struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon_species"`
		Rate int `json:"rate"`
	} `json:"pokemon_encounters"`
}
