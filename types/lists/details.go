package lists

import "encoding/json"

func UnmarshalDetails(data []byte) (Details, error) {
	var r Details
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Details) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Details struct {
	CreatedBy     string `json:"created_by"`
	Description   string `json:"description"`
	FavoriteCount int    `json:"favorite_count"`
	ID            string `json:"id"`
	Items         []Item `json:"items"`
	ItemCount     int    `json:"item_count"`
	ISO639_1      string `json:"iso_639_1"`
	Name          string `json:"name"`
	PosterPath    string `json:"poster_path"`
}

type Item struct {
	PosterPath       string      `json:"poster_path"`
	Adult            bool        `json:"adult"`
	Overview         string      `json:"overview"`
	ReleaseDate      string      `json:"release_date"`
	OriginalTitle    string      `json:"original_title"`
	GenreIDS         []int       `json:"genre_ids"`
	ID               int         `json:"id"`
	MediaType        string      `json:"media_type"`
	OriginalLanguage string      `json:"original_language"`
	Title            string      `json:"title"`
	BackdropPath     interface{} `json:"backdrop_path"`
	Popularity       float64     `json:"popularity"`
	VoteCount        int         `json:"vote_count"`
	Video            bool        `json:"video"`
	VoteAverage      float64     `json:"vote_average"`
}
