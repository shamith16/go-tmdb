package movies

import "encoding/json"

func UnmarshalNowPlaying(data []byte) (NowPlaying, error) {
	var r NowPlaying
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *NowPlaying) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type NowPlaying struct {
	Page         int      `json:"page"`
	Results      []Result `json:"results"`
	Dates        Dates    `json:"dates"`
	TotalPages   int      `json:"total_pages"`
	TotalResults int      `json:"total_results"`
}
