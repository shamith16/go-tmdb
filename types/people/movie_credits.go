package people

import "encoding/json"

func UnmarshalMovieCredits(data []byte) (MovieCredits, error) {
	var r MovieCredits
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MovieCredits) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type MovieCredits struct {
	Cast []Cast `json:"cast"`
}
