package tv_episodes

import "encoding/json"

func UnmarshalDetails(data []byte) (Details, error) {
	var r Details
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Details) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Details struct {
	AirDate        string      `json:"air_date"`
	Crew           []Crew      `json:"crew"`
	EpisodeNumber  int         `json:"episode_number"`
	GuestStars     []GuestStar `json:"guest_stars"`
	Name           string      `json:"name"`
	Overview       string      `json:"overview"`
	ID             int         `json:"id"`
	ProductionCode string      `json:"production_code"`
	SeasonNumber   int         `json:"season_number"`
	StillPath      string      `json:"still_path"`
	VoteAverage    float64     `json:"vote_average"`
	VoteCount      int         `json:"vote_count"`
}

type Crew struct {
	ID          int     `json:"id"`
	CreditID    string  `json:"credit_id"`
	Name        string  `json:"name"`
	Department  string  `json:"department"`
	Job         string  `json:"job"`
	ProfilePath *string `json:"profile_path"`
}

type GuestStar struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	CreditID    string `json:"credit_id"`
	Character   string `json:"character"`
	Order       int    `json:"order"`
	ProfilePath string `json:"profile_path"`
}
