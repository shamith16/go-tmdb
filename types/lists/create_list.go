package lists

import "encoding/json"

func UnmarshalCreateList(data []byte) (CreateList, error) {
	var r CreateList
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateList) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type CreateList struct {
	StatusMessage string `json:"status_message"`
	Success       bool   `json:"success"`
	StatusCode    int    `json:"status_code"`
	ListID        int    `json:"list_id"`
}
