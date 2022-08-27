package movies

import "encoding/json"

func UnmarshalLists(data []byte) (Lists, error) {
	var r Lists
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Lists) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Lists struct {
	ID           int      `json:"id"`
	Page         int      `json:"page"`
	Results      []Result `json:"results"`
	TotalPages   int      `json:"total_pages"`
	TotalResults int      `json:"total_results"`
}
