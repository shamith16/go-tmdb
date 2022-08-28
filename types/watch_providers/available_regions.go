package watch_providers

import "encoding/json"

func UnmarshalAvailableRegions(data []byte) (AvailableRegions, error) {
	var r AvailableRegions
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AvailableRegions) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AvailableRegions struct {
	Results []Result `json:"results"`
}
