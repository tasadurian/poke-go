package types

// ContestEffect ...
type ContestEffect struct {
	Appeal        int `json:"appeal"`
	EffectEntries []struct {
		Effect   string `json:"effect"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"effect_entries"`
	FlavorTextEntries []struct {
		FlavorText string `json:"flavor_text"`
		Language   struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"flavor_text_entries"`
	ID  int `json:"id"`
	Jam int `json:"jam"`
}
