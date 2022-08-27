package account

import "encoding/json"

func UnmarshalCreatedList(data []byte) (CreatedList, error) {
	var r CreatedList
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreatedList) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type CreatedList struct {
	Page         int      `json:"page"`
	Results      []Result `json:"results"`
	TotalPages   int      `json:"total_pages"`
	TotalResults int      `json:"total_results"`
}
