package networks

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
	ID    int    `json:"id"`
	Logos []Logo `json:"logos"`
}

type Logo struct {
	AspectRatio float64 `json:"aspect_ratio"`
	FilePath    string  `json:"file_path"`
	Height      int     `json:"height"`
	ID          string  `json:"id"`
	FileType    string  `json:"file_type"`
	VoteAverage int     `json:"vote_average"`
	VoteCount   int     `json:"vote_count"`
	Width       int     `json:"width"`
}
