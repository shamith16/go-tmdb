package movies

import "encoding/json"

func UnmarshalKeywords(data []byte) (Keywords, error) {
	var r Keywords
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Keywords) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Keywords struct {
	ID       int       `json:"id"`
	Keywords []Keyword `json:"keywords"`
}

type Keyword struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
