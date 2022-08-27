package find

import "encoding/json"

func UnmarshalID(data []byte) (ID, error) {
	var r ID
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ID) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ID struct {
	MovieResults     []MovieResult `json:"movie_results"`
	PersonResults    []interface{} `json:"person_results"`
	TvResults        []interface{} `json:"tv_results"`
	TvEpisodeResults []interface{} `json:"tv_episode_results"`
	TvSeasonResults  []interface{} `json:"tv_season_results"`
}

type MovieResult struct {
	Adult            bool        `json:"adult"`
	BackdropPath     interface{} `json:"backdrop_path"`
	GenreIDS         []int       `json:"genre_ids"`
	ID               int         `json:"id"`
	OriginalLanguage string      `json:"original_language"`
	OriginalTitle    string      `json:"original_title"`
	Overview         string      `json:"overview"`
	ReleaseDate      string      `json:"release_date"`
	PosterPath       interface{} `json:"poster_path"`
	Popularity       float64     `json:"popularity"`
	Title            string      `json:"title"`
	Video            bool        `json:"video"`
	VoteAverage      float64     `json:"vote_average"`
	VoteCount        int         `json:"vote_count"`
}
