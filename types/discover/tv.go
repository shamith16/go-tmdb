package discover

import "encoding/json"

func UnmarshalTv(data []byte) (Tv, error) {
	var r Tv
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Tv) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Tv struct {
	Page         int      `json:"page"`
	Results      []Result `json:"results"`
	TotalResults int      `json:"total_results"`
	TotalPages   int      `json:"total_pages"`
}
