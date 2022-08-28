package people

import "encoding/json"

func UnmarshalTranslations(data []byte) (Translations, error) {
	var r Translations
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Translations) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Translations struct {
	Translations []Translation `json:"translations"`
}

type Translation struct {
	ISO639_1    string `json:"iso_639_1"`
	ISO3166_1   string `json:"iso_3166_1"`
	Name        string `json:"name"`
	Data        Data   `json:"data"`
	EnglishName string `json:"english_name"`
}

type Data struct {
	Biography string `json:"biography"`
}
