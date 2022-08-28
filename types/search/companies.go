package search

import "encoding/json"

func UnmarshalCompanies(data []byte) (Companies, error) {
	var r Companies
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Companies) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Companies struct {
	Page         int      `json:"page"`
	Results      []Result `json:"results"`
	TotalPages   int      `json:"total_pages"`
	TotalResults int      `json:"total_results"`
}
