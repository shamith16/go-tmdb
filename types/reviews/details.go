package reviews

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
	ID            string        `json:"id"`
	Author        string        `json:"author"`
	AuthorDetails AuthorDetails `json:"author_details"`
	Content       string        `json:"content"`
	CreatedAt     string        `json:"created_at"`
	ISO639_1      string        `json:"iso_639_1"`
	MediaID       int           `json:"media_id"`
	MediaTitle    string        `json:"media_title"`
	MediaType     string        `json:"media_type"`
	UpdatedAt     string        `json:"updated_at"`
	URL           string        `json:"url"`
}

type AuthorDetails struct {
	Name       string `json:"name"`
	Username   string `json:"username"`
	AvatarPath string `json:"avatar_path"`
	Rating     int    `json:"rating"`
}
