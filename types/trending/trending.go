package trending

import "encoding/json"

func UnmarshalTrending(data []byte) (Trending, error) {
	var r Trending
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Trending) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Trending struct {
	Page         int      `json:"page"`
	Results      []Result `json:"results"`
	TotalPages   int      `json:"total_pages"`
	TotalResults int      `json:"total_results"`
}

type Result struct {
	Adult            *bool    `json:"adult,omitempty"`
	BackdropPath     string   `json:"backdrop_path"`
	GenreIDS         []int    `json:"genre_ids"`
	ID               int      `json:"id"`
	OriginalLanguage string   `json:"original_language"`
	OriginalTitle    *string  `json:"original_title,omitempty"`
	Overview         string   `json:"overview"`
	PosterPath       string   `json:"poster_path"`
	ReleaseDate      *string  `json:"release_date,omitempty"`
	Title            *string  `json:"title,omitempty"`
	Video            *bool    `json:"video,omitempty"`
	VoteAverage      float64  `json:"vote_average"`
	VoteCount        int      `json:"vote_count"`
	Popularity       float64  `json:"popularity"`
	FirstAirDate     *string  `json:"first_air_date,omitempty"`
	Name             *string  `json:"name,omitempty"`
	OriginCountry    []string `json:"origin_country,omitempty"`
	OriginalName     *string  `json:"original_name,omitempty"`
}
