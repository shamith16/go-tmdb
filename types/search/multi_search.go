package search

import "encoding/json"

func UnmarshalMultiSearch(data []byte) (MultiSearch, error) {
	var r MultiSearch
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MultiSearch) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type MultiSearch struct {
	Page         int      `json:"page"`
	Results      []Result `json:"results"`
	TotalResults int      `json:"total_results"`
	TotalPages   int      `json:"total_pages"`
}
