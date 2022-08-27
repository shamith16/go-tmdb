package guest_sessions

type Result struct {
	AirDate          string      `json:"air_date,omitempty"`
	EpisodeNumber    int         `json:"episode_number,omitempty"`
	ID               int         `json:"id,omitempty"`
	Name             string      `json:"name,omitempty"`
	Overview         string      `json:"overview,omitempty"`
	ProductionCode   *string     `json:"production_code,omitempty"`
	SeasonNumber     int         `json:"season_number,omitempty"`
	ShowID           int         `json:"show_id,omitempty"`
	StillPath        *string     `json:"still_path,omitempty"`
	VoteAverage      float64     `json:"vote_average,omitempty"`
	VoteCount        int         `json:"vote_count,omitempty"`
	Rating           float64     `json:"rating,omitempty"`
	BackdropPath     string      `json:"backdrop_path,omitempty"`
	FirstAirDate     string      `json:"first_air_date,omitempty"`
	GenreIDS         []int       `json:"genre_ids,omitempty"`
	OriginalLanguage interface{} `json:"original_language,omitempty"`
	OriginalName     string      `json:"original_name,omitempty"`
	OriginCountry    []string    `json:"origin_country,omitempty"`
	PosterPath       string      `json:"poster_path,omitempty"`
	Popularity       float64     `json:"popularity,omitempty"`
	Adult            bool        `json:"adult,omitempty"`
	OriginalTitle    string      `json:"original_title,omitempty"`
	ReleaseDate      string      `json:"release_date,omitempty"`
	Title            string      `json:"title,omitempty"`
	Video            bool        `json:"video,omitempty"`
}
