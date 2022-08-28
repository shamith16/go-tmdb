package tv

import "encoding/json"

func UnmarshalContentRatings(data []byte) (ContentRatings, error) {
	var r ContentRatings
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ContentRatings) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ContentRatings struct {
	Results []Result `json:"results"`
	ID      int      `json:"id"`
}
