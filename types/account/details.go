package account

import "encoding/json"

func UnmarshalDetails(data []byte) (Details, error) {
	var r Details
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Details) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Details struct {
	Avatar       Avatar `json:"avatar"`
	ID           int    `json:"id"`
	ISO639_1     string `json:"iso_639_1"`
	ISO3166_1    string `json:"iso_3166_1"`
	Name         string `json:"name"`
	IncludeAdult bool   `json:"include_adult"`
	Username     string `json:"username"`
}

type Avatar struct {
	Gravatar Gravatar `json:"gravatar"`
}

type Gravatar struct {
	Hash string `json:"hash"`
}
