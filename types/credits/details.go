package credits

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
	CreditType string `json:"credit_type"`
	Department string `json:"department"`
	Job        string `json:"job"`
	Media      Media  `json:"media"`
	MediaType  string `json:"media_type"`
	ID         string `json:"id"`
	Person     Person `json:"person"`
}

type Media struct {
	ID           int           `json:"id"`
	Name         string        `json:"name"`
	OriginalName string        `json:"original_name"`
	Character    string        `json:"character"`
	Episodes     []interface{} `json:"episodes"`
	Seasons      []Season      `json:"seasons"`
}

type Season struct {
	AirDate      string `json:"air_date"`
	PosterPath   string `json:"poster_path"`
	SeasonNumber int    `json:"season_number"`
}

type Person struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}
