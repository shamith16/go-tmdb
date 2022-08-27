package movies

type Genre struct {
	ID   int    `json:"id"`
	Name string `json:"name,omitempty"`
}

type Dates struct {
	Maximum string `json:"maximum,omitempty"`
	Minimum string `json:"minimum,omitempty"`
}

type Result struct {
	PosterPath       string        `json:"poster_path,omitempty"`
	Adult            bool          `json:"adult,omitempty"`
	Overview         string        `json:"overview,omitempty"`
	ReleaseDate      string        `json:"release_date,omitempty"`
	GenreIDS         []int         `json:"genre_ids,omitempty"`
	ID               int           `json:"id,omitempty"`
	OriginalTitle    string        `json:"original_title,omitempty"`
	OriginalLanguage string        `json:"original_language,omitempty"`
	Title            string        `json:"title,omitempty"`
	BackdropPath     string        `json:"backdrop_path,omitempty"`
	Popularity       float64       `json:"popularity,omitempty"`
	VoteCount        int           `json:"vote_count,omitempty"`
	Video            bool          `json:"video,omitempty"`
	VoteAverage      float64       `json:"vote_average,omitempty"`
	Description      string        `json:"description,omitempty"`
	FavoriteCount    int           `json:"favorite_count,omitempty"`
	ItemCount        int           `json:"item_count,omitempty"`
	ISO639_1         string        `json:"iso_639_1,omitempty"`
	ListType         string        `json:"list_type,omitempty"`
	Name             string        `json:"name,omitempty"`
	Author           string        `json:"author,omitempty"`
	AuthorDetails    AuthorDetails `json:"author_details,omitempty"`
	Content          string        `json:"content,omitempty"`
	CreatedAt        string        `json:"created_at,omitempty"`
	UpdatedAt        string        `json:"updated_at,omitempty"`
	URL              string        `json:"url,omitempty"`
	ISO3166_1        string        `json:"iso_3166_1,omitempty"`
	ReleaseDates     []ReleaseDate `json:"release_dates,omitempty"`
	Key              string        `json:"key,omitempty"`
	Site             string        `json:"site,omitempty"`
	Size             int           `json:"size,omitempty"`
	Type             string        `json:"type,omitempty"`
	Official         bool          `json:"official,omitempty"`
	PublishedAt      string        `json:"published_at,omitempty"`
}
