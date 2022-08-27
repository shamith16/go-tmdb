package collections

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
	ID           int         `json:"id"`
	Name         string      `json:"name"`
	Overview     string      `json:"overview"`
	PosterPath   interface{} `json:"poster_path"`
	BackdropPath string      `json:"backdrop_path"`
	Parts        []Part      `json:"parts"`
}

type Part struct {
	Adult            bool        `json:"adult"`
	BackdropPath     interface{} `json:"backdrop_path"`
	GenreIDS         []int       `json:"genre_ids"`
	ID               int         `json:"id"`
	OriginalLanguage string      `json:"original_language"`
	OriginalTitle    string      `json:"original_title"`
	Overview         string      `json:"overview"`
	ReleaseDate      string      `json:"release_date"`
	PosterPath       string      `json:"poster_path"`
	Popularity       float64     `json:"popularity"`
	Title            string      `json:"title"`
	Video            bool        `json:"video"`
	VoteAverage      float64     `json:"vote_average"`
	VoteCount        int         `json:"vote_count"`
}
