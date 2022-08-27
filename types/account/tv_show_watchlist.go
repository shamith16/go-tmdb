package account

import "encoding/json"

func UnmarshalTvShowWatchlist(data []byte) (TvShowWatchlist, error) {
	var r TvShowWatchlist
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TvShowWatchlist) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type TvShowWatchlist struct {
	Page         int      `json:"page"`
	Results      []Result `json:"results"`
	TotalPages   int      `json:"total_pages"`
	TotalResults int      `json:"total_results"`
}
