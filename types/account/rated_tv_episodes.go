package account

import "encoding/json"

func UnmarshalRatedTvEpisodes(data []byte) (RatedTvEpisodes, error) {
	var r RatedTvEpisodes
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *RatedTvEpisodes) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type RatedTvEpisodes struct {
	Page         int      `json:"page"`
	Results      []Result `json:"results"`
	TotalPages   int      `json:"total_pages"`
	TotalResults int      `json:"total_results"`
}
