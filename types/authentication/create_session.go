package authentication

import "encoding/json"

func UnmarshalCreateSession(data []byte) (CreateSession, error) {
	var r CreateSession
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateSession) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type CreateSession struct {
	Success   bool   `json:"success"`
	SessionID string `json:"session_id"`
}
