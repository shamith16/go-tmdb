package authentication

import "encoding/json"

func UnmarshalCreateRequestToken(data []byte) (CreateRequestToken, error) {
	var r CreateRequestToken
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateRequestToken) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type CreateRequestToken struct {
	Success      bool   `json:"success"`
	ExpiresAt    string `json:"expires_at"`
	RequestToken string `json:"request_token"`
}
