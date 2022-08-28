package tv_seasons

import "encoding/json"

func UnmarshalVideos(data []byte) (Videos, error) {
	var r Videos
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Videos) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Videos struct {
	ID      int      `json:"id"`
	Results []Result `json:"results"`
}
