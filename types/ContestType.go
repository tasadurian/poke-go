package poke

// ContestType ...
type ContestType struct {
	BerryFlavor struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"berry_flavor"`
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Names []struct {
		Color    string `json:"color"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
}
