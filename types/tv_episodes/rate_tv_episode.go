package tv_episodes

import "encoding/json"

func UnmarshalRateTvEpisode(data []byte) (RateTvEpisode, error) {
	var r RateTvEpisode
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *RateTvEpisode) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type RateTvEpisode struct {
	StatusCode    int    `json:"status_code"`
	StatusMessage string `json:"status_message"`
}
