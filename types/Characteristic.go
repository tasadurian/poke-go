package types

// Characteristic ...
type Characteristic struct {
	Descriptions []struct {
		Description string `json:"description"`
		Language    struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"descriptions"`
	GeneModulo  int `json:"gene_modulo"`
	HighestStat struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"highest_stat"`
	ID             int   `json:"id"`
	PossibleValues []int `json:"possible_values"`
}
