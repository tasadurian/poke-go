package pokego

type PokeathlonStat struct {
	AffectingNatures struct {
		Decrease []struct {
			MaxChange int `json:"max_change"`
			Nature    struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"nature"`
		} `json:"decrease"`
		Increase []struct {
			MaxChange int `json:"max_change"`
			Nature    struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"nature"`
		} `json:"increase"`
	} `json:"affecting_natures"`
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
