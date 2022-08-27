package account

import "encoding/json"

func UnmarshalMarkFavourite(data []byte) (MarkFavourite, error) {
	var r MarkFavourite
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MarkFavourite) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type MarkFavourite struct {
	StatusCode    int    `json:"status_code"`
	StatusMessage string `json:"status_message"`
}
