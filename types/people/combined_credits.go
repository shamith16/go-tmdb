package people

import "encoding/json"

func UnmarshalCombinedCredits(data []byte) (CombinedCredits, error) {
	var r CombinedCredits
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CombinedCredits) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type CombinedCredits struct {
	Cast []Cast `json:"cast"`
	Crew []Cast `json:"crew"`
	ID   int    `json:"id"`
}
