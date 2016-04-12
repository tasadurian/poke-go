package pokego

// ItemCategory
type ItemCategory struct {
	ID    int `json:"id"`
	Items []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"items"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	Pocket struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"pocket"`
}
