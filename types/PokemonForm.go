package pokego

// PokemonForm ...
type PokemonForm struct {
	FormName     string `json:"form_name"`
	FormOrder    int    `json:"form_order"`
	ID           int    `json:"id"`
	IsBattleOnly bool   `json:"is_battle_only"`
	IsDefault    bool   `json:"is_default"`
	IsMega       bool   `json:"is_mega"`
	Name         string `json:"name"`
	Order        int    `json:"order"`
	Pokemon      struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"pokemon"`
	Sprites struct {
		BackDefault  string `json:"back_default"`
		BackShiny    string `json:"back_shiny"`
		FrontDefault string `json:"front_default"`
		FrontShiny   string `json:"front_shiny"`
	} `json:"sprites"`
	VersionGroup struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"version_group"`
}
