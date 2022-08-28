package tv

import "encoding/json"

func UnmarshalAccountStates(data []byte) (AccountStates, error) {
	var r AccountStates
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AccountStates) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AccountStates struct {
	ID        int   `json:"id"`
	Favorite  bool  `json:"favorite"`
	Rated     Rated `json:"rated"`
	Watchlist bool  `json:"watchlist"`
}

type Rated struct {
	Value int `json:"value"`
}
