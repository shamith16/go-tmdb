package tv

type Result struct {
	ISO639_1         string        `json:"iso_639_1"`
	ISO3166_1        string        `json:"iso_3166_1"`
	Name             string        `json:"name"`
	Key              string        `json:"key"`
	Site             string        `json:"site"`
	Size             int           `json:"size"`
	Type             string        `json:"type"`
	Official         bool          `json:"official"`
	PublishedAt      string        `json:"published_at"`
	ID               string        `json:"id"`
	PosterPath       string        `json:"poster_path"`
	Popularity       float64       `json:"popularity"`
	BackdropPath     *string       `json:"backdrop_path"`
	VoteAverage      float64       `json:"vote_average"`
	Overview         string        `json:"overview"`
	FirstAirDate     string        `json:"first_air_date"`
	OriginCountry    []string      `json:"origin_country"`
	GenreIDS         []int         `json:"genre_ids"`
	OriginalLanguage string        `json:"original_language"`
	VoteCount        int           `json:"vote_count"`
	OriginalName     string        `json:"original_name"`
	EpisodeNumber    int           `json:"episode_number"`
	SeasonNumber     int           `json:"season_number"`
	Author           string        `json:"author"`
	AuthorDetails    AuthorDetails `json:"author_details"`
	Content          string        `json:"content"`
	CreatedAt        string        `json:"created_at"`
	UpdatedAt        string        `json:"updated_at"`
	URL              string        `json:"url"`
	Description      string        `json:"description"`
	EpisodeCount     int           `json:"episode_count"`
	GroupCount       int           `json:"group_count"`
	Network          *Network      `json:"network"`
	Rating           string        `json:"rating"`
}

type Genre struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Season struct {
	AirDate      string      `json:"air_date"`
	EpisodeCount int         `json:"episode_count"`
	ID           int         `json:"id"`
	PosterPath   interface{} `json:"poster_path"`
	SeasonNumber int         `json:"season_number"`
	Name         string      `json:"name"`
	Overview     string      `json:"overview"`
}

type Network struct {
	ID            int     `json:"id"`
	LogoPath      *string `json:"logo_path"`
	Name          string  `json:"name"`
	OriginCountry string  `json:"origin_country"`
}

type Cast struct {
	Adult              bool    `json:"adult"`
	Gender             int     `json:"gender"`
	ID                 int     `json:"id"`
	KnownForDepartment string  `json:"known_for_department"`
	Name               string  `json:"name"`
	OriginalName       string  `json:"original_name"`
	Popularity         float64 `json:"popularity"`
	ProfilePath        *string `json:"profile_path"`
	Character          *string `json:"character,omitempty"`
	CreditID           string  `json:"credit_id"`
	Order              *int    `json:"order,omitempty"`
	Department         *string `json:"department,omitempty"`
	Job                *string `json:"job,omitempty"`
	Roles              []Role  `json:"roles,omitempty"`
	TotalEpisodeCount  int     `json:"total_episode_count"`
	Jobs               []Job   `json:"jobs,omitempty"`
}
