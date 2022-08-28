package tv

import "encoding/json"

func UnmarshalKeywords(data []byte) (Keywords, error) {
	var r Keywords
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Keywords) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Keywords struct {
	ID      int      `json:"id"`
	Results []Result `json:"results"`
}
