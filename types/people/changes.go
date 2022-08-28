package people

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
	ID            string        `json:"id"`
	Action        string        `json:"action"`
	Time          string        `json:"time"`
	OriginalValue OriginalValue `json:"original_value"`
}

type OriginalValue struct {
	Profile Profile `json:"profile"`
}
