package search

import "encoding/json"

func UnmarshalMovies(data []byte) (Movies, error) {
	var r Movies
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Movies) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Movies struct {
	Page         int      `json:"page"`
	Results      []Result `json:"results"`
	TotalResults int      `json:"total_results"`
	TotalPages   int      `json:"total_pages"`
}
