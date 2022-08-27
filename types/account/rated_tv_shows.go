package account

import "encoding/json"

func UnmarshalRatedTvShows(data []byte) (RatedTvShows, error) {
	var r RatedTvShows
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *RatedTvShows) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type RatedTvShows struct {
	Page         int      `json:"page"`
	Results      []Result `json:"results"`
	TotalPages   int      `json:"total_pages"`
	TotalResults int      `json:"total_results"`
}
