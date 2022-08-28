package tv

import "encoding/json"

func UnmarshalTopRated(data []byte) (TopRated, error) {
	var r TopRated
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TopRated) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type TopRated struct {
	Page         int      `json:"page"`
	Results      []Result `json:"results"`
	TotalResults int      `json:"total_results"`
	TotalPages   int      `json:"total_pages"`
}
