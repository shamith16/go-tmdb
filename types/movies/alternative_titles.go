package movies

import "encoding/json"

func UnmarshalAlternativeTitles(data []byte) (AlternativeTitles, error) {
	var r AlternativeTitles
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AlternativeTitles) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AlternativeTitles struct {
	ID     int     `json:"id"`
	Titles []Title `json:"titles"`
}

type Title struct {
	ISO3166_1 string `json:"iso_3166_1"`
	Title     string `json:"title"`
	Type      string `json:"type"`
}
