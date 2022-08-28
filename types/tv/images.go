package tv

import "encoding/json"

func UnmarshalImages(data []byte) (Images, error) {
	var r Images
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Images) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Images struct {
	Backdrops []Backdrop `json:"backdrops"`
	ID        int        `json:"id"`
	Posters   []Backdrop `json:"posters"`
}

type Backdrop struct {
	AspectRatio float64 `json:"aspect_ratio"`
	FilePath    string  `json:"file_path"`
	Height      int     `json:"height"`
	ISO639_1    *string `json:"iso_639_1"`
	VoteAverage float64 `json:"vote_average"`
	VoteCount   int     `json:"vote_count"`
	Width       int     `json:"width"`
}
