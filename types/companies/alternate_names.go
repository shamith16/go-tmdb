package companies

import "encoding/json"

func UnmarshalAlternateNames(data []byte) (AlternateNames, error) {
	var r AlternateNames
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AlternateNames) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AlternateNames struct {
	ID      int      `json:"id"`
	Results []Result `json:"results"`
}

type Result struct {
	Name string `json:"name"`
	Type string `json:"type"`
}
