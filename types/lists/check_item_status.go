package lists

import "encoding/json"

func UnmarshalCheckItemStatus(data []byte) (CheckItemStatus, error) {
	var r CheckItemStatus
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CheckItemStatus) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type CheckItemStatus struct {
	ID          string `json:"id"`
	ItemPresent bool   `json:"item_present"`
}
