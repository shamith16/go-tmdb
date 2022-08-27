package genres

import "encoding/json"

func UnmarshalTvlist(data []byte) (Tvlist, error) {
	var r Tvlist
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Tvlist) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Tvlist struct {
	Genres []Genre `json:"genres"`
}
