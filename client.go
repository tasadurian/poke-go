package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/thetommytwitch/poke-go/types"
	"io/ioutil"
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

// Client ...
type Client struct{}

// GetPokeData ...
// make this concurent and add a timeout param
func (c *Client) request(value string, items []string) ([][]byte, error) {

	if items == nil {
		return nil, errors.New("Err: No parameters found, need at least one item in slice")
	}

	var response [][]byte

	for _, resource := range items {
		res, err := http.Get(endpoint + value + "/" + resource + "/")
		if err != nil {
			return nil, err
		}

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		response = append(response, body)
	}

	return response, nil
}

// GetAbility ...
func (c *Client) GetAbility(items []string) ([]types.Ability, error) {
	abilities := []types.Ability{}
	responses, err := c.request("ability", items)
	if err != nil {
		return nil, err
	}
	for _, response := range responses {
		a := types.Ability{}
		err := json.Unmarshal(response, &a)
		if err != nil {
			return nil, err
		}
		abilities = append(abilities, a)
	}
	return abilities, nil
}

func main() {
	client := new(Client)
	params := []string{"1"}
	ability, err := client.GetAbility(params)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ability)
}
