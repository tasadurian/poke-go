package clients

// BerryClient ...
type BerryClient struct {
	Name string
	ID   int
}

// GetBerry ...
func (bc *BerryClient) GetBerry() (berry Berry, err error) {
	url := endpoint + "/berry/" + bc.Name
	if err = EndpointRequest(url, &berry); err != nil {
		return Berry{}, err
	}
	return berry, nil
}
