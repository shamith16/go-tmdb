package tv_episodes

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

type Result struct {
	ISO639_1    string `json:"iso_639_1"`
	ISO3166_1   string `json:"iso_3166_1"`
	Name        string `json:"name"`
	Key         string `json:"key"`
	PublishedAt string `json:"published_at"`
	Site        string `json:"site"`
	Size        int    `json:"size"`
	Type        string `json:"type"`
	Official    bool   `json:"official"`
	ID          string `json:"id"`
}
