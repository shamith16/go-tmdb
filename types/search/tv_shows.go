package search

import "encoding/json"

func UnmarshalTvShows(data []byte) (TvShows, error) {
	var r TvShows
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TvShows) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type TvShows struct {
	Page         int      `json:"page"`
	Results      []Result `json:"results"`
	TotalResults int      `json:"total_results"`
	TotalPages   int      `json:"total_pages"`
}
