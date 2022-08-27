package companies

import "encoding/json"

func UnmarshalDetails(data []byte) (Details, error) {
	var r Details
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Details) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Details struct {
	Description   string      `json:"description"`
	Headquarters  string      `json:"headquarters"`
	Homepage      string      `json:"homepage"`
	ID            int         `json:"id"`
	LogoPath      string      `json:"logo_path"`
	Name          string      `json:"name"`
	OriginCountry string      `json:"origin_country"`
	ParentCompany interface{} `json:"parent_company"`
}
