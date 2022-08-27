package configuration

import "encoding/json"

func UnmarshalConfiguration(data []byte) (Configuration, error) {
	var r Configuration
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Configuration) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Configuration struct {
	Images     Images   `json:"images"`
	ChangeKeys []string `json:"change_keys"`
}

type Images struct {
	BaseURL       string   `json:"base_url"`
	SecureBaseURL string   `json:"secure_base_url"`
	BackdropSizes []string `json:"backdrop_sizes"`
	LogoSizes     []string `json:"logo_sizes"`
	PosterSizes   []string `json:"poster_sizes"`
	ProfileSizes  []string `json:"profile_sizes"`
	StillSizes    []string `json:"still_sizes"`
}
