package movies

import "encoding/json"

func UnmarshalSimilarMovies(data []byte) (SimilarMovies, error) {
	var r SimilarMovies
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SimilarMovies) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type SimilarMovies struct {
	Page         int      `json:"page"`
	Results      []Result `json:"results"`
	TotalPages   int      `json:"total_pages"`
	TotalResults int      `json:"total_results"`
}
