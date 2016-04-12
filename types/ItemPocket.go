package pokego

// ItemPocket
type ItemPocket struct {
	Categories []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"categories"`
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
}
