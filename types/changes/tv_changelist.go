package changes

import "encoding/json"

func UnmarshalTvChangelist(data []byte) (TvChangelist, error) {
	var r TvChangelist
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TvChangelist) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type TvChangelist struct {
	Results      []Result `json:"results"`
	Page         int      `json:"page"`
	TotalPages   int      `json:"total_pages"`
	TotalResults int      `json:"total_results"`
}
