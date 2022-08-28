package tv_seasons

import "encoding/json"

func UnmarshalAccountStates(data []byte) (AccountStates, error) {
	var r AccountStates
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AccountStates) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AccountStates struct {
	ID      int      `json:"id"`
	Results []Result `json:"results"`
}

type RatedClass struct {
	Value int `json:"value"`
}

type RatedUnion struct {
	Bool       *bool
	RatedClass *RatedClass
}

func (x *RatedUnion) UnmarshalJSON(data []byte) error {
	x.RatedClass = nil
	var c RatedClass
	object, err := unmarshalUnion(data, nil, nil, &x.Bool, nil, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.RatedClass = &c
	}
	return nil
}

func (x *RatedUnion) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, x.Bool, nil, false, nil, x.RatedClass != nil, x.RatedClass, false, nil, false, nil, false)
}
