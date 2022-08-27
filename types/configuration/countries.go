package configuration

import "encoding/json"

type Countries []Country

func UnmarshalCountries(data []byte) (Countries, error) {
	var r Countries
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Countries) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Country struct {
	ISO639_1    string `json:"iso_639_1"`
	EnglishName string `json:"english_name"`
	Name        string `json:"name"`
}
