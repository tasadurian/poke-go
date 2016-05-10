package clients

import (
	types "github.com/thetommytwitch/poke-go/types"
)

// BerryClient ...
type BerryClient struct {
	Name string
	ID   int
}

// GetBerry ...
func (bc BerryClient) GetBerry() (berry types.Berry, err error) {
	if bc.ID == nil {
		url := endpoint + "/berry/" + bc.Name
	}
	if bc.Name == nil {
		url := endpoint + "/berry/" + bc.ID
	}
	if err = EndpointRequest(url, &berry); err != nil {
		return Berry{}, err
	}
	return berry, nil
}
