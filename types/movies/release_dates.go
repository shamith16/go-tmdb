package movies

import "encoding/json"

func UnmarshalReleaseDates(data []byte) (ReleaseDates, error) {
	var r ReleaseDates
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ReleaseDates) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ReleaseDates struct {
	ID      int      `json:"id"`
	Results []Result `json:"results"`
}

type ReleaseDate struct {
	Certification string  `json:"certification"`
	ISO639_1      string  `json:"iso_639_1"`
	ReleaseDate   string  `json:"release_date"`
	Type          int     `json:"type"`
	Note          *string `json:"note,omitempty"`
}
