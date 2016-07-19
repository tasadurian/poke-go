package main

import (
	"encoding/json"
	"errors"
	"github.com/thetommytwitch/poke-go/types"
	"net/http"
)

const (
	endpoint = "http://pokeapi.co/api/v2/"
)

var dataMap = map[string]interface{}{
	"ability":                   types.Ability{},
	"berry":                     types.Berry{},
	"berry-firmness":            types.BerryFirmness{},
	"berry-flavor":              types.BerryFlavor{},
	"characteristic":            types.Characteristic{},
	"contest-effect":            types.ContestEffect{},
	"contest-type":              types.ContestType{},
	"egg-group":                 types.EggGroup{},
	"encounter-condition":       types.EncounterCondition{},
	"encounter-condition-value": types.EncounterConditionValue{},
	"encounter-method":          types.EncounterMethod{},
	"evolution-chain":           types.EvolutionChain{},
	"evolution-trigger":         types.EvolutionTrigger{},
	"gender":                    types.Gender{},
	"generation":                types.Generation{},
	"growth-rate":               types.GrowthRate{},
	"item":                      types.Item{},
	"item-attribute":            types.ItemAttribute{},
	"item-category":             types.ItemCategory{},
	"item-fling-effect":         types.ItemFlingEffect{},
	"item-pocket":               types.ItemPocket{},
	"language":                  types.Language{},
	"location":                  types.Location{},
	"location-area":             types.LocationArea{},
	"move":                      types.Move{},
	"move-ailment":              types.MoveAilment{},
	"move-battle-style":         types.MoveBattleStyle{},
	"move-category":             types.MoveCategory{},
	"move-damage-class":         types.MoveDamageClass{},
	"move-learn-method":         types.MoveLearnMethod{},
	"move-target":               types.MoveTarget{},
	"nature":                    types.Nature{},
	"pal-park-area":             types.PalParkArea{},
	"pokeathlon-stat":           types.PokeathlonStat{},
	"pokedex":                   types.Pokedex{},
	"pokemon":                   types.Pokemon{},
	"pokemon-color":             types.PokemonColor{},
	"pokemon-form":              types.PokemonForm{},
	"pokemon-habitat":           types.PokemonHabitat{},
	"pokemon-shape":             types.PokemonShape{},
	"pokemon-species":           types.PokemonSpecies{},
	"region":                    types.Region{},
	"stat":                      types.Stat{},
	"super-contest-effect":      types.SuperContestEffect{},
	"type":                      types.Type{},
	"version":                   types.Version{},
	"version-group":             types.VersionGroup{},
}

// GetPokeData ...
func GetPokeData(value string, pokemon string) (interface{}, error) {

	res, err := http.Get(endpoint + value + "/" + pokemon + "/")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	for key, v := range dataMap {
		if key == value {
			decoder.Decode(&v)
			return v, nil
		}
	}
	return nil, errors.New("Err: value not found.")
}
