package discover

type Result struct {
	PosterPath       string      `json:"poster_path,omitempty"`
	Popularity       float64     `json:"popularity,omitempty"`
	ID               int         `json:"id,omitempty"`
	BackdropPath     interface{} `json:"backdrop_path,omitempty"`
	VoteAverage      float64     `json:"vote_average,omitempty"`
	Overview         string      `json:"overview,omitempty"`
	FirstAirDate     string      `json:"first_air_date,omitempty"`
	OriginCountry    []string    `json:"origin_country,omitempty"`
	GenreIDS         []int       `json:"genre_ids,omitempty"`
	OriginalLanguage string      `json:"original_language,omitempty"`
	VoteCount        int         `json:"vote_count,omitempty"`
	Name             string      `json:"name,omitempty"`
	OriginalName     string      `json:"original_name,omitempty"`
	Adult            bool        `json:"adult,omitempty"`
	ReleaseDate      string      `json:"release_date,omitempty"`
	OriginalTitle    string      `json:"original_title,omitempty"`
	Title            string      `json:"title,omitempty"`
	Video            bool        `json:"video,omitempty"`
}
