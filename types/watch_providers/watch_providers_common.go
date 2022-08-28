package watch_providers

type Result struct {
	DisplayPriority int    `json:"display_priority,omitempty"`
	LogoPath        string `json:"logo_path,omitempty"`
	ProviderName    string `json:"provider_name,omitempty"`
	ProviderID      int    `json:"provider_id,omitempty"`
	ISO3166_1       string `json:"iso_3166_1,omitempty"`
	EnglishName     string `json:"english_name,omitempty"`
	NativeName      string `json:"native_name,omitempty"`
}
