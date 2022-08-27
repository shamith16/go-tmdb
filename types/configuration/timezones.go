package configuration

import "encoding/json"

type Timezones []Timezone

func UnmarshalTimezones(data []byte) (Timezones, error) {
	var r Timezones
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Timezones) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Timezone struct {
	ISO3166_1 string   `json:"iso_3166_1"`
	Zones     []string `json:"zones"`
}
