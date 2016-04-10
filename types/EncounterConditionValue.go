package pokego

// EncounterConditionValue
type EncounterConditionValue struct {
	Condition struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"condition"`
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
