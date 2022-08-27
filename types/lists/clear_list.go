package lists

import "encoding/json"

func UnmarshalClearList(data []byte) (ClearList, error) {
	var r ClearList
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ClearList) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ClearList struct {
	StatusCode    int    `json:"status_code"`
	StatusMessage string `json:"status_message"`
}
