package tv_episodes

import "encoding/json"

func UnmarshalChanges(data []byte) (Changes, error) {
	var r Changes
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Changes) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Changes struct {
	Changes []Change `json:"changes"`
}

type Change struct {
	Key   string `json:"key"`
	Items []Item `json:"items"`
}

type Item struct {
	ID            string  `json:"id"`
	Action        string  `json:"action"`
	Time          string  `json:"time"`
	Value         string  `json:"value"`
	ISO639_1      *string `json:"iso_639_1,omitempty"`
	OriginalValue *string `json:"original_value,omitempty"`
}
