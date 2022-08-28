package tv_seasons

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
	ID           string    `json:"_id"`
	AirDate      string    `json:"air_date"`
	Episodes     []Episode `json:"episodes"`
	Name         string    `json:"name"`
	Overview     string    `json:"overview"`
	DetailsID    int       `json:"id"`
	PosterPath   string    `json:"poster_path"`
	SeasonNumber int       `json:"season_number"`
}

type Episode struct {
	AirDate        string  `json:"air_date"`
	EpisodeNumber  int     `json:"episode_number"`
	Crew           []Crew  `json:"crew"`
	GuestStars     []Crew  `json:"guest_stars"`
	ID             int     `json:"id"`
	Name           string  `json:"name"`
	Overview       string  `json:"overview"`
	ProductionCode string  `json:"production_code"`
	SeasonNumber   int     `json:"season_number"`
	StillPath      string  `json:"still_path"`
	VoteAverage    float64 `json:"vote_average"`
	VoteCount      int     `json:"vote_count"`
}

type Crew struct {
	Department         *string `json:"department,omitempty"`
	Job                *string `json:"job,omitempty"`
	CreditID           string  `json:"credit_id"`
	Adult              bool    `json:"adult"`
	Gender             int     `json:"gender"`
	ID                 int     `json:"id"`
	KnownForDepartment string  `json:"known_for_department"`
	Name               string  `json:"name"`
	OriginalName       string  `json:"original_name"`
	Popularity         float64 `json:"popularity"`
	ProfilePath        *string `json:"profile_path"`
	Order              *int    `json:"order,omitempty"`
	Character          *string `json:"character,omitempty"`
}
