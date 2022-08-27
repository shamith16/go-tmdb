package account

import "encoding/json"

func UnmarshalFavouriteTvShows(data []byte) (FavouriteTvShows, error) {
	var r FavouriteTvShows
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FavouriteTvShows) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type FavouriteTvShows struct {
	Page         int      `json:"page"`
	Results      []Result `json:"results"`
	TotalPages   int      `json:"total_pages"`
	TotalResults int      `json:"total_results"`
}
