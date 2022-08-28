package people

import "encoding/json"

func UnmarshalTaggedImages(data []byte) (TaggedImages, error) {
	var r TaggedImages
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TaggedImages) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type TaggedImages struct {
	ID           int      `json:"id"`
	Page         int      `json:"page"`
	Results      []Result `json:"results"`
	TotalPages   int      `json:"total_pages"`
	TotalResults int      `json:"total_results"`
}

type Media struct {
	Adult            bool    `json:"adult"`
	BackdropPath     string  `json:"backdrop_path"`
	GenreIDS         []int   `json:"genre_ids"`
	ID               string  `json:"_id"`
	MediaID          int     `json:"id"`
	OriginalLanguage string  `json:"original_language"`
	OriginalTitle    string  `json:"original_title"`
	Overview         string  `json:"overview"`
	ReleaseDate      string  `json:"release_date"`
	PosterPath       string  `json:"poster_path"`
	Popularity       float64 `json:"popularity"`
	Title            string  `json:"title"`
	Video            bool    `json:"video"`
	VoteAverage      float64 `json:"vote_average"`
	VoteCount        int     `json:"vote_count"`
}
