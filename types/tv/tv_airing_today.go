package tv

import "encoding/json"

func UnmarshalTvAiringToday(data []byte) (TvAiringToday, error) {
	var r TvAiringToday
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TvAiringToday) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type TvAiringToday struct {
	Page         int      `json:"page"`
	Results      []Result `json:"results"`
	TotalResults int      `json:"total_results"`
	TotalPages   int      `json:"total_pages"`
}
