package tv

import "encoding/json"

func UnmarshalScreenedTheatrically(data []byte) (ScreenedTheatrically, error) {
	var r ScreenedTheatrically
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ScreenedTheatrically) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ScreenedTheatrically struct {
	ID      int      `json:"id"`
	Results []Result `json:"results"`
}
