package testapp

import "encoding/json"

type status struct {
	Status string
}

func (sta status) Encode() ([]byte, string, error) {
	data, err := json.Marshal(sta)
	return data, "application/json", err
}
