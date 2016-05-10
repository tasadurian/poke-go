package main

import (
	"encoding/json"
	"net/http"
)

const (
	endpoint = "http://pokeapi.co/api/v2"
)

// EndpointRequest ...
func EndpointRequest(url string, value interface{}) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	return decoder.Decode(&value)
}
