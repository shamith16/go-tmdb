package tv

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
	BackdropPath        interface{}   `json:"backdrop_path"`
	CreatedBy           []interface{} `json:"created_by"`
	EpisodeRunTime      []int         `json:"episode_run_time"`
	FirstAirDate        string        `json:"first_air_date"`
	Genres              []Genre       `json:"genres"`
	Homepage            string        `json:"homepage"`
	ID                  int           `json:"id"`
	InProduction        bool          `json:"in_production"`
	Languages           []string      `json:"languages"`
	LastAirDate         string        `json:"last_air_date"`
	Name                string        `json:"name"`
	Networks            []Genre       `json:"networks"`
	NumberOfEpisodes    int           `json:"number_of_episodes"`
	NumberOfSeasons     int           `json:"number_of_seasons"`
	OriginCountry       []string      `json:"origin_country"`
	OriginalLanguage    string        `json:"original_language"`
	OriginalName        string        `json:"original_name"`
	Overview            interface{}   `json:"overview"`
	Popularity          int           `json:"popularity"`
	PosterPath          interface{}   `json:"poster_path"`
	ProductionCompanies []interface{} `json:"production_companies"`
	Seasons             []Season      `json:"seasons"`
	Status              string        `json:"status"`
	Type                string        `json:"type"`
	VoteAverage         int           `json:"vote_average"`
	VoteCount           int           `json:"vote_count"`
}
