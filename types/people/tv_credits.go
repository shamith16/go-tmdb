package people

import "encoding/json"

func UnmarshalTvCredits(data []byte) (TvCredits, error) {
	var r TvCredits
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TvCredits) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type TvCredits struct {
	Cast []Cast `json:"cast"`
}
