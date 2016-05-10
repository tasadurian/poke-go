package types

// BerryFlavor ...
type BerryFlavor struct {
	Berries []struct {
		Berry struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"berry"`
		Potency int `json:"potency"`
	} `json:"berries"`
	ContestType struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"contest_type"`
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
