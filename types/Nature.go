package types

// Nature ...
type Nature struct {
	DecreasedStat struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"decreased_stat"`
	HatesFlavor struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"hates_flavor"`
	ID            int `json:"id"`
	IncreasedStat struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"increased_stat"`
	LikesFlavor struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"likes_flavor"`
	MoveBattleStylePreferences []struct {
		HighHpPreference int `json:"high_hp_preference"`
		LowHpPreference  int `json:"low_hp_preference"`
		MoveBattleStyle  struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"move_battle_style"`
	} `json:"move_battle_style_preferences"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokeathlonStatChanges []struct {
		MaxChange      int `json:"max_change"`
		PokeathlonStat struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokeathlon_stat"`
	} `json:"pokeathlon_stat_changes"`
}
