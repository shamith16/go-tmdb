package authentication

import "encoding/json"

func UnmarshalCreateSessionLogin(data []byte) (CreateSessionLogin, error) {
	var r CreateSessionLogin
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateSessionLogin) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type CreateSessionLogin struct {
	Success      bool   `json:"success"`
	ExpiresAt    string `json:"expires_at"`
	RequestToken string `json:"request_token"`
}
