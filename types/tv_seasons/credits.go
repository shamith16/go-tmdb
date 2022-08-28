package tv_seasons

import "encoding/json"

func UnmarshalCredits(data []byte) (Credits, error) {
	var r Credits
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Credits) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Credits struct {
	Cast []Cast `json:"cast"`
	Crew []Cast `json:"crew"`
	ID   int    `json:"id"`
}
