package keywords

import "encoding/json"

func UnmarshalMovies(data []byte) (Movies, error) {
	var r Movies
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Movies) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Movies struct {
	ID           int      `json:"id"`
	Page         int      `json:"page"`
	Results      []Result `json:"results"`
	TotalPages   int      `json:"total_pages"`
	TotalResults int      `json:"total_results"`
}

type Result struct {
	Adult            bool        `json:"adult"`
	BackdropPath     interface{} `json:"backdrop_path"`
	GenreIDS         []int       `json:"genre_ids"`
	ID               int         `json:"id"`
	OriginalLanguage string      `json:"original_language"`
	OriginalTitle    string      `json:"original_title"`
	Overview         string      `json:"overview"`
	ReleaseDate      string      `json:"release_date"`
	PosterPath       string      `json:"poster_path"`
	Popularity       float64     `json:"popularity"`
	Title            string      `json:"title"`
	Video            bool        `json:"video"`
	VoteAverage      float64     `json:"vote_average"`
	VoteCount        int         `json:"vote_count"`
}
