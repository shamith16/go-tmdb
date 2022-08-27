package account

import "encoding/json"

func UnmarshalMovieWatchlist(data []byte) (MovieWatchlist, error) {
	var r MovieWatchlist
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MovieWatchlist) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type MovieWatchlist struct {
	Page         int      `json:"page"`
	Results      []Result `json:"results"`
	TotalPages   int      `json:"total_pages"`
	TotalResults int      `json:"total_results"`
}
