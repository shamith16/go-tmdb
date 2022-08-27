package account

import "encoding/json"

func UnmarshalRatedMovies(data []byte) (RatedMovies, error) {
	var r RatedMovies
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *RatedMovies) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type RatedMovies struct {
	Page         int      `json:"page"`
	Results      []Result `json:"results"`
	TotalPages   int      `json:"total_pages"`
	TotalResults int      `json:"total_results"`
}
