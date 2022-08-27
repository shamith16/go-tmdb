package authentication

import "encoding/json"

func UnmarshalCreateSessionV4(data []byte) (CreateSessionV4, error) {
	var r CreateSessionV4
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateSessionV4) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type CreateSessionV4 struct {
	Success   bool   `json:"success"`
	SessionID string `json:"session_id"`
}
