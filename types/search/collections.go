package search

import "encoding/json"

func UnmarshalCollections(data []byte) (Collections, error) {
	var r Collections
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Collections) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Collections struct {
	Page         int      `json:"page"`
	Results      []Result `json:"results"`
	TotalPages   int      `json:"total_pages"`
	TotalResults int      `json:"total_results"`
}
