package people

import "encoding/json"

func UnmarshalImages(data []byte) (Images, error) {
	var r Images
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Images) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Images struct {
	ID       int       `json:"id"`
	Profiles []Profile `json:"profiles"`
}
