package watch_providers

import "encoding/json"

func UnmarshalTvProviders(data []byte) (TvProviders, error) {
	var r TvProviders
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TvProviders) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type TvProviders struct {
	Results []Result `json:"results"`
}
