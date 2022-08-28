package search

import "encoding/json"

func UnmarshalPeople(data []byte) (People, error) {
	var r People
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *People) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type People struct {
	Page         int      `json:"page"`
	Results      []Result `json:"results"`
	TotalResults int      `json:"total_results"`
	TotalPages   int      `json:"total_pages"`
}

type KnownFor struct {
	PosterPath       string   `json:"poster_path"`
	Adult            *bool    `json:"adult,omitempty"`
	Overview         string   `json:"overview"`
	ReleaseDate      *string  `json:"release_date,omitempty"`
	OriginalTitle    *string  `json:"original_title,omitempty"`
	GenreIDS         []int    `json:"genre_ids"`
	ID               int      `json:"id"`
	MediaType        string   `json:"media_type"`
	OriginalLanguage string   `json:"original_language"`
	Title            *string  `json:"title,omitempty"`
	BackdropPath     *string  `json:"backdrop_path"`
	Popularity       float64  `json:"popularity"`
	VoteCount        int      `json:"vote_count"`
	Video            *bool    `json:"video,omitempty"`
	VoteAverage      float64  `json:"vote_average"`
	FirstAirDate     *string  `json:"first_air_date,omitempty"`
	OriginCountry    []string `json:"origin_country,omitempty"`
	Name             *string  `json:"name,omitempty"`
	OriginalName     *string  `json:"original_name,omitempty"`
}
