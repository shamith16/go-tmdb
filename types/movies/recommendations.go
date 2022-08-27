package movies

import "encoding/json"

func UnmarshalRecommendations(data []byte) (Recommendations, error) {
	var r Recommendations
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Recommendations) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Recommendations struct {
	Page         int      `json:"page"`
	Results      []Result `json:"results"`
	TotalPages   int      `json:"total_pages"`
	TotalResults int      `json:"total_results"`
}
