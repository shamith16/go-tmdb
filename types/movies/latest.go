package movies

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
	Adult               bool          `json:"adult"`
	BackdropPath        interface{}   `json:"backdrop_path"`
	BelongsToCollection interface{}   `json:"belongs_to_collection"`
	Budget              int           `json:"budget"`
	Genres              []Genre       `json:"genres"`
	Homepage            string        `json:"homepage"`
	ID                  int           `json:"id"`
	ImdbID              string        `json:"imdb_id"`
	OriginalLanguage    string        `json:"original_language"`
	OriginalTitle       string        `json:"original_title"`
	Overview            string        `json:"overview"`
	Popularity          int           `json:"popularity"`
	PosterPath          string        `json:"poster_path"`
	ProductionCompanies []interface{} `json:"production_companies"`
	ProductionCountries []interface{} `json:"production_countries"`
	ReleaseDate         string        `json:"release_date"`
	Revenue             int           `json:"revenue"`
	Runtime             int           `json:"runtime"`
	SpokenLanguages     []interface{} `json:"spoken_languages"`
	Status              string        `json:"status"`
	Tagline             string        `json:"tagline"`
	Title               string        `json:"title"`
	Video               bool          `json:"video"`
	VoteAverage         int           `json:"vote_average"`
	VoteCount           int           `json:"vote_count"`
}
