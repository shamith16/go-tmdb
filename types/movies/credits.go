package movies

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
	ID   int    `json:"id"`
	Cast []Cast `json:"cast"`
}

type Cast struct {
	Adult              bool    `json:"adult"`
	Gender             int     `json:"gender"`
	ID                 int     `json:"id"`
	KnownForDepartment string  `json:"known_for_department"`
	Name               string  `json:"name"`
	OriginalName       string  `json:"original_name"`
	Popularity         float64 `json:"popularity"`
	ProfilePath        *string `json:"profile_path"`
	CastID             int     `json:"cast_id"`
	Character          string  `json:"character"`
	CreditID           string  `json:"credit_id"`
	Order              int     `json:"order"`
}
