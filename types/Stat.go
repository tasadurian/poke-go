package types

// Stat ...
type Stat struct {
	AffectingMoves struct {
		Decrease []struct {
			Change int `json:"change"`
			Move   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"move"`
		} `json:"decrease"`
		Increase []struct {
			Change int `json:"change"`
			Move   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"move"`
		} `json:"increase"`
	} `json:"affecting_moves"`
	AffectingNatures struct {
		Decrease []struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"decrease"`
		Increase []struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"increase"`
	} `json:"affecting_natures"`
	Characteristics []struct {
		URL string `json:"url"`
	} `json:"characteristics"`
	GameIndex       int  `json:"game_index"`
	ID              int  `json:"id"`
	IsBattleOnly    bool `json:"is_battle_only"`
	MoveDamageClass struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"move_damage_class"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
}
