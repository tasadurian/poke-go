package types

// ItemFlingEffect ...
type ItemFlingEffect struct {
	EffectEntries []struct {
		Effect   string `json:"effect"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"effect_entries"`
	ID    int `json:"id"`
	Items []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"items"`
	Name string `json:"name"`
}
