package changes

import "encoding/json"

func UnmarshalMovieChangelist(data []byte) (MovieChangelist, error) {
	var r MovieChangelist
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MovieChangelist) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type MovieChangelist struct {
	Results      []Result `json:"results"`
	Page         int      `json:"page"`
	TotalPages   int      `json:"total_pages"`
	TotalResults int      `json:"total_results"`
}
