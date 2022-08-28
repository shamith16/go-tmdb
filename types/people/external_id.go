package people

import "encoding/json"

func UnmarshalExternalID(data []byte) (ExternalID, error) {
	var r ExternalID
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ExternalID) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ExternalID struct {
	ImdbID      string `json:"imdb_id"`
	FacebookID  string `json:"facebook_id"`
	FreebaseMid string `json:"freebase_mid"`
	FreebaseID  string `json:"freebase_id"`
	TvrageID    int    `json:"tvrage_id"`
	TwitterID   string `json:"twitter_id"`
	ID          int    `json:"id"`
}
