package networks

import "encoding/json"

func UnmarshalAlternativeNames(data []byte) (AlternativeNames, error) {
	var r AlternativeNames
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AlternativeNames) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AlternativeNames struct {
	ID      int      `json:"id"`
	Results []Result `json:"results"`
}

type Result struct {
	Name string `json:"name"`
	Type string `json:"type"`
}
