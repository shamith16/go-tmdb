package authentication

import "encoding/json"

func UnmarshalDeleteSession(data []byte) (DeleteSession, error) {
	var r DeleteSession
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteSession) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type DeleteSession struct {
	Success bool `json:"success"`
}
