package tv_episode_groups

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
	Description  string  `json:"description"`
	EpisodeCount int     `json:"episode_count"`
	GroupCount   int     `json:"group_count"`
	Groups       []Group `json:"groups"`
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Network      Network `json:"network"`
	Type         int     `json:"type"`
}

type Group struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Order    int       `json:"order"`
	Episodes []Episode `json:"episodes"`
	Locked   bool      `json:"locked"`
}

type Episode struct {
	AirDate        string      `json:"air_date"`
	EpisodeNumber  int         `json:"episode_number"`
	ID             int         `json:"id"`
	Name           string      `json:"name"`
	Overview       string      `json:"overview"`
	ProductionCode interface{} `json:"production_code"`
	SeasonNumber   int         `json:"season_number"`
	ShowID         int         `json:"show_id"`
	StillPath      *string     `json:"still_path"`
	VoteAverage    int         `json:"vote_average"`
	VoteCount      int         `json:"vote_count"`
	Order          int         `json:"order"`
}

type Network struct {
	ID            int    `json:"id"`
	LogoPath      string `json:"logo_path"`
	Name          string `json:"name"`
	OriginCountry string `json:"origin_country"`
}
