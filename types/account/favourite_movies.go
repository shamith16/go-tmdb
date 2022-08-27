package account

import "encoding/json"

func UnmarshalFavouriteMovies(data []byte) (FavouriteMovies, error) {
	var r FavouriteMovies
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FavouriteMovies) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type FavouriteMovies struct {
	Page         int      `json:"page"`
	Results      []Result `json:"results"`
	TotalPages   int      `json:"total_pages"`
	TotalResults int      `json:"total_results"`
}
