package tv_episodes

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
	FreebaseMid string `json:"freebase_mid"`
	FreebaseID  string `json:"freebase_id"`
	TvdbID      int    `json:"tvdb_id"`
	TvrageID    int    `json:"tvrage_id"`
	ID          int    `json:"id"`
}
