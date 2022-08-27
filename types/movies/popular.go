package movies

import "encoding/json"

func UnmarshalPopular(data []byte) (Popular, error) {
	var r Popular
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Popular) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Popular struct {
	Page         int      `json:"page"`
	Results      []Result `json:"results"`
	TotalResults int      `json:"total_results"`
	TotalPages   int      `json:"total_pages"`
}
