package tv_seasons

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
	FreebaseMid string      `json:"freebase_mid"`
	FreebaseID  interface{} `json:"freebase_id"`
	TvdbID      int         `json:"tvdb_id"`
	TvrageID    interface{} `json:"tvrage_id"`
	ID          int         `json:"id"`
}
