package lists

import "encoding/json"

func UnmarshalRemoveMovie(data []byte) (RemoveMovie, error) {
	var r RemoveMovie
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *RemoveMovie) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type RemoveMovie struct {
	StatusCode    int    `json:"status_code"`
	StatusMessage string `json:"status_message"`
}
