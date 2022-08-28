package tv_episodes

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
	Cast       []Cast `json:"cast"`
	Crew       []Cast `json:"crew"`
	GuestStars []Cast `json:"guest_stars"`
	ID         int    `json:"id"`
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
	Character          *string `json:"character,omitempty"`
	CreditID           string  `json:"credit_id"`
	Order              *int    `json:"order,omitempty"`
	Department         *string `json:"department,omitempty"`
	Job                *string `json:"job,omitempty"`
	CharacterName      *string `json:"character_name,omitempty"`
}
