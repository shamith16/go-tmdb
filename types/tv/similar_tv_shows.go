package tv

import "encoding/json"

func UnmarshalSimilarTvShows(data []byte) (SimilarTvShows, error) {
	var r SimilarTvShows
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SimilarTvShows) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type SimilarTvShows struct {
	Page         int      `json:"page"`
	Results      []Result `json:"results"`
	TotalPages   int      `json:"total_pages"`
	TotalResults int      `json:"total_results"`
}
