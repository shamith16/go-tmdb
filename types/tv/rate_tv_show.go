package tv

import "encoding/json"

func UnmarshalRateTvShow(data []byte) (RateTvShow, error) {
	var r RateTvShow
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *RateTvShow) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type RateTvShow struct {
	StatusCode    int    `json:"status_code"`
	StatusMessage string `json:"status_message"`
}
