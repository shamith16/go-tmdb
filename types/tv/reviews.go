package tv

import "encoding/json"

func UnmarshalReviews(data []byte) (Reviews, error) {
	var r Reviews
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Reviews) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Reviews struct {
	ID           int      `json:"id"`
	Page         int      `json:"page"`
	Results      []Result `json:"results"`
	TotalPages   int      `json:"total_pages"`
	TotalResults int      `json:"total_results"`
}

type AuthorDetails struct {
	Name       string  `json:"name"`
	Username   string  `json:"username"`
	AvatarPath *string `json:"avatar_path"`
	Rating     int     `json:"rating"`
}
