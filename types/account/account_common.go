package account

type Result struct {
	Adult            bool        `json:"adult,omitempty"`
	BackdropPath     interface{} `json:"backdrop_path,omitempty"`
	GenreIDS         []int       `json:"genre_ids,omitempty"`
	ID               int         `json:"id,omitempty"`
	OriginalLanguage string      `json:"original_language,omitempty"`
	OriginalTitle    string      `json:"original_title,omitempty"`
	Overview         string      `json:"overview,omitempty"`
	ReleaseDate      string      `json:"release_date,omitempty"`
	PosterPath       interface{} `json:"poster_path,omitempty"`
	Popularity       float64     `json:"popularity,omitempty"`
	Title            string      `json:"title,omitempty"`
	Video            bool        `json:"video,omitempty"`
	VoteAverage      float64     `json:"vote_average,omitempty"`
	VoteCount        int         `json:"vote_count,omitempty"`
	Description      string      `json:"description,omitempty"`
	FavoriteCount    int         `json:"favorite_count,omitempty"`
	ItemCount        int         `json:"item_count,omitempty"`
	ISO639_1         string      `json:"iso_639_1,omitempty"`
	ListType         string      `json:"list_type,omitempty"`
	FirstAirDate     string      `json:"first_air_date,omitempty"`
	OriginalName     string      `json:"original_name,omitempty"`
	OriginCountry    []string    `json:"origin_country,omitempty"`
	Name             string      `json:"name,omitempty"`
	Rating           int         `json:"rating,omitempty"`
	AirDate          string      `json:"air_date,omitempty"`
	EpisodeNumber    int         `json:"episode_number,omitempty"`
	ProductionCode   *string     `json:"production_code,omitempty"`
	SeasonNumber     int         `json:"season_number,omitempty"`
	ShowID           int         `json:"show_id,omitempty"`
	StillPath        string      `json:"still_path,omitempty"`
}
