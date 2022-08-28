package watch_providers

import "encoding/json"

func UnmarshalMovieProviders(data []byte) (MovieProviders, error) {
	var r MovieProviders
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MovieProviders) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type MovieProviders struct {
	Results []Result `json:"results"`
}
