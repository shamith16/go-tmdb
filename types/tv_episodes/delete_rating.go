package tv_episodes

import "encoding/json"

func UnmarshalDeleteRating(data []byte) (DeleteRating, error) {
	var r DeleteRating
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteRating) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type DeleteRating struct {
	StatusCode    int    `json:"status_code"`
	StatusMessage string `json:"status_message"`
}
