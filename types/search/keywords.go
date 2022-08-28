package search

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
	Page         int      `json:"page"`
	Results      []Result `json:"results"`
	TotalPages   int      `json:"total_pages"`
	TotalResults int      `json:"total_results"`
}
