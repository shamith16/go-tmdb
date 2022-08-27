package authentication

import "encoding/json"

func UnmarshalCreateGuestSession(data []byte) (CreateGuestSession, error) {
	var r CreateGuestSession
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateGuestSession) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type CreateGuestSession struct {
	Success        bool   `json:"success"`
	GuestSessionID string `json:"guest_session_id"`
	ExpiresAt      string `json:"expires_at"`
}
