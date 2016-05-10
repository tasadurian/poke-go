package types

// VersionGroup ...
type VersionGroup struct {
	Generation struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"generation"`
	ID               int `json:"id"`
	MoveLearnMethods []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"move_learn_methods"`
	Name      string `json:"name"`
	Order     int    `json:"order"`
	Pokedexes []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"pokedexes"`
	Regions []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"regions"`
	Versions []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"versions"`
}
