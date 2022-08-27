package changes

import "encoding/json"

func UnmarshalPersonChangelist(data []byte) (PersonChangelist, error) {
	var r PersonChangelist
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PersonChangelist) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type PersonChangelist struct {
	Results      []Result `json:"results"`
	Page         int      `json:"page"`
	TotalPages   int      `json:"total_pages"`
	TotalResults int      `json:"total_results"`
}
