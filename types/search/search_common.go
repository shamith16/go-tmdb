package search

type Result struct {
	PosterPath       string     `json:"poster_path,omitempty"`
	Popularity       float64    `json:"popularity,omitempty"`
	ID               int        `json:"id,omitempty"`
	BackdropPath     string     `json:"backdrop_path,omitempty"`
	VoteAverage      float64    `json:"vote_average,omitempty"`
	Overview         string     `json:"overview,omitempty"`
	FirstAirDate     string     `json:"first_air_date,omitempty"`
	OriginCountry    []string   `json:"origin_country,omitempty"`
	GenreIDS         []int      `json:"genre_ids,omitempty"`
	OriginalLanguage string     `json:"original_language,omitempty"`
	VoteCount        int        `json:"vote_count,omitempty"`
	Name             string     `json:"name,omitempty"`
	OriginalName     string     `json:"original_name,omitempty"`
	ProfilePath      *string    `json:"profile_path,omitempty"`
	Adult            bool       `json:"adult,omitempty"`
	KnownFor         []KnownFor `json:"known_for,omitempty"`
	MediaType        string     `json:"media_type,omitempty"`
	ReleaseDate      *string    `json:"release_date,omitempty,omitempty"`
	OriginalTitle    *string    `json:"original_title,omitempty,omitempty"`
	Title            *string    `json:"title,omitempty,omitempty"`
	Video            *bool      `json:"video,omitempty,omitempty"`
	LogoPath         *string    `json:"logo_path,omitempty"`
}
