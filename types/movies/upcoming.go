package movies

import "encoding/json"

func UnmarshalUpcoming(data []byte) (Upcoming, error) {
	var r Upcoming
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Upcoming) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Upcoming struct {
	Page         int      `json:"page"`
	Results      []Result `json:"results"`
	Dates        Dates    `json:"dates"`
	TotalPages   int      `json:"total_pages"`
	TotalResults int      `json:"total_results"`
}
