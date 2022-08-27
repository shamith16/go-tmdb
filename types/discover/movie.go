package discover

import "encoding/json"

func UnmarshalMovie(data []byte) (Movie, error) {
	var r Movie
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Movie) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Movie struct {
	Page         int      `json:"page"`
	Results      []Result `json:"results"`
	TotalResults int      `json:"total_results"`
	TotalPages   int      `json:"total_pages"`
}

type OriginalLanguage string

const (
	En OriginalLanguage = "en"
)
