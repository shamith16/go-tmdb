package people

type Profile struct {
	AspectRatio float64 `json:"aspect_ratio,omitempty"`
	FilePath    string  `json:"file_path,omitempty"`
	Height      int     `json:"height,omitempty"`
	ISO639_1    string  `json:"iso_639_1,omitempty"`
	VoteAverage float64 `json:"vote_average,omitempty"`
	VoteCount   int     `json:"vote_count,omitempty"`
	Width       int     `json:"width,omitempty"`
}

type Result struct {
	AspectRatio float64    `json:"aspect_ratio,omitempty"`
	FilePath    string     `json:"file_path,omitempty"`
	Height      int        `json:"height,omitempty"`
	ID          string     `json:"id,omitempty"`
	ISO639_1    string     `json:"iso_639_1,omitempty"`
	VoteAverage float64    `json:"vote_average,omitempty"`
	VoteCount   int        `json:"vote_count,omitempty"`
	Width       int        `json:"width,omitempty"`
	ImageType   string     `json:"image_type,omitempty"`
	Media       Media      `json:"media,omitempty"`
	MediaType   string     `json:"media_type,omitempty"`
	ProfilePath string     `json:"profile_path,omitempty"`
	Adult       bool       `json:"adult,omitempty"`
	KnownFor    []KnownFor `json:"known_for,omitempty"`
	Name        string     `json:"name,omitempty"`
	Popularity  float64    `json:"popularity,omitempty"`
}

type Cast struct {
	CreditID         string   `json:"credit_id,omitempty"`
	OriginalName     string   `json:"original_name,omitempty"`
	ID               int      `json:"id,omitempty"`
	GenreIDS         []int    `json:"genre_ids,omitempty"`
	Character        string   `json:"character,omitempty"`
	Name             string   `json:"name,omitempty"`
	PosterPath       string   `json:"poster_path,omitempty"`
	VoteCount        int      `json:"vote_count,omitempty"`
	VoteAverage      float64  `json:"vote_average,omitempty"`
	Popularity       float64  `json:"popularity,omitempty"`
	EpisodeCount     int      `json:"episode_count,omitempty"`
	OriginalLanguage string   `json:"original_language,omitempty"`
	FirstAirDate     string   `json:"first_air_date,omitempty"`
	BackdropPath     string   `json:"backdrop_path,omitempty"`
	Overview         string   `json:"overview,omitempty"`
	OriginCountry    []string `json:"origin_country,omitempty"`
	ReleaseDate      string   `json:"release_date,omitempty"`
	Video            bool     `json:"video,omitempty"`
	Adult            bool     `json:"adult,omitempty"`
	Title            string   `json:"title,omitempty"`
	OriginalTitle    string   `json:"original_title,omitempty"`
	MediaType        string   `json:"media_type,omitempty"`
	Department       *string  `json:"department,omitempty,omitempty"`
	Job              *string  `json:"job,omitempty,omitempty"`
}
