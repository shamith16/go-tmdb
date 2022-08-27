package movies

import "encoding/json"

func UnmarshalRateMovie(data []byte) (RateMovie, error) {
	var r RateMovie
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *RateMovie) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type RateMovie struct {
	StatusCode    int    `json:"status_code"`
	StatusMessage string `json:"status_message"`
}
