package lists

import "encoding/json"

func UnmarshalDeleteList(data []byte) (DeleteList, error) {
	var r DeleteList
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteList) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type DeleteList struct {
	StatusCode    int    `json:"status_code"`
	StatusMessage string `json:"status_message"`
}
