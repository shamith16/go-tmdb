package account

import "encoding/json"

func UnmarshalAddWatchlist(data []byte) (AddWatchlist, error) {
	var r AddWatchlist
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AddWatchlist) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AddWatchlist struct {
	StatusCode    int    `json:"status_code"`
	StatusMessage string `json:"status_message"`
}
