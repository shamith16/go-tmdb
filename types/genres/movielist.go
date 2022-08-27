package genres

import "encoding/json"

func UnmarshalMovielist(data []byte) (Movielist, error) {
	var r Movielist
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Movielist) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Movielist struct {
	Genres []Genre `json:"genres"`
}
