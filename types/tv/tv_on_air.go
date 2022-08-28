package tv

import "encoding/json"

func UnmarshalTvOnAir(data []byte) (TvOnAir, error) {
	var r TvOnAir
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TvOnAir) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type TvOnAir struct {
	Page         int      `json:"page"`
	Results      []Result `json:"results"`
	TotalResults int      `json:"total_results"`
	TotalPages   int      `json:"total_pages"`
}
