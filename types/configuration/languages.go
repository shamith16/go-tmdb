package configuration

import "encoding/json"

type Languages []Language

func UnmarshalLanguages(data []byte) (Languages, error) {
	var r Languages
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Languages) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Language struct {
	ISO639_1    string `json:"iso_639_1"`
	EnglishName string `json:"english_name"`
	Name        string `json:"name"`
}
