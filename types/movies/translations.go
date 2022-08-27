package movies

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
	ID           int           `json:"id"`
	Translations []Translation `json:"translations"`
}

type Translation struct {
	ISO3166_1   string `json:"iso_3166_1"`
	ISO639_1    string `json:"iso_639_1"`
	Name        string `json:"name"`
	EnglishName string `json:"english_name"`
	Data        Data   `json:"data"`
}

type Data struct {
	Title    string `json:"title"`
	Overview string `json:"overview"`
	Homepage string `json:"homepage"`
}
