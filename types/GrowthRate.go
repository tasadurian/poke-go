package types

// GrowthRate ...
type GrowthRate struct {
	Descriptions []struct {
		Description string `json:"description"`
		Language    struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"descriptions"`
	Formula string `json:"formula"`
	ID      int    `json:"id"`
	Levels  []struct {
		Experience int `json:"experience"`
		Level      int `json:"level"`
	} `json:"levels"`
	Name           string `json:"name"`
	PokemonSpecies []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"pokemon_species"`
}
