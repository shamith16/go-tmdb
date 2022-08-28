package people

import "encoding/json"

func UnmarshalLatest(data []byte) (Latest, error) {
	var r Latest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Latest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Latest struct {
	Adult              bool          `json:"adult"`
	AlsoKnownAs        []interface{} `json:"also_known_as"`
	Biography          string        `json:"biography"`
	Birthday           interface{}   `json:"birthday"`
	Deathday           interface{}   `json:"deathday"`
	Gender             int           `json:"gender"`
	Homepage           interface{}   `json:"homepage"`
	ID                 int           `json:"id"`
	ImdbID             interface{}   `json:"imdb_id"`
	KnownForDepartment string        `json:"known_for_department"`
	Name               string        `json:"name"`
	PlaceOfBirth       interface{}   `json:"place_of_birth"`
	Popularity         int           `json:"popularity"`
	ProfilePath        interface{}   `json:"profile_path"`
}
