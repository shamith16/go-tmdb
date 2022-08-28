package tv

import "encoding/json"

func UnmarshalEpisodeGroups(data []byte) (EpisodeGroups, error) {
	var r EpisodeGroups
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *EpisodeGroups) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type EpisodeGroups struct {
	Results []Result `json:"results"`
	ID      int      `json:"id"`
}
