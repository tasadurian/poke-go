package types

// Berries are small fruits that can provide HP and status condition restoration,
// stat enhancement, and even damage negation when eaten by Pokémon.
type Berry struct {
	Firmness struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"firmness"`
	Flavors []struct {
		Flavor struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"flavor"`
		Potency int `json:"potency"`
	} `json:"flavors"`
	GrowthTime int `json:"growth_time"`
	ID         int `json:"id"`
	Item       struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"item"`
	MaxHarvest       int    `json:"max_harvest"`
	Name             string `json:"name"`
	NaturalGiftPower int    `json:"natural_gift_power"`
	NaturalGiftType  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"natural_gift_type"`
	Size        int `json:"size"`
	Smoothness  int `json:"smoothness"`
	SoilDryness int `json:"soil_dryness"`
}
