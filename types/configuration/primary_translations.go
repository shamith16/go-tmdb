package configuration

import "encoding/json"

type PrimaryTranslations []string

func UnmarshalPrimaryTranslations(data []byte) (PrimaryTranslations, error) {
	var r PrimaryTranslations
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PrimaryTranslations) Marshal() ([]byte, error) {
	return json.Marshal(r)
}
