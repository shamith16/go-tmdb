package people

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
	Birthday           string      `json:"birthday"`
	KnownForDepartment string      `json:"known_for_department"`
	Deathday           interface{} `json:"deathday"`
	ID                 int         `json:"id"`
	Name               string      `json:"name"`
	AlsoKnownAs        []string    `json:"also_known_as"`
	Gender             int         `json:"gender"`
	Biography          string      `json:"biography"`
	Popularity         float64     `json:"popularity"`
	PlaceOfBirth       string      `json:"place_of_birth"`
	ProfilePath        string      `json:"profile_path"`
	Adult              bool        `json:"adult"`
	ImdbID             string      `json:"imdb_id"`
	Homepage           interface{} `json:"homepage"`
}
