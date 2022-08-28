package tv_seasons

import "bytes"
import "errors"
import "encoding/json"

func UnmarshalChanges(data []byte) (Changes, error) {
	var r Changes
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Changes) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Changes struct {
	Changes []Change `json:"changes"`
}

type Change struct {
	Key   string `json:"key"`
	Items []Item `json:"items"`
}

type Item struct {
	ID            string      `json:"id"`
	Action        string      `json:"action"`
	Time          string      `json:"time"`
	Value         *ValueUnion `json:"value"`
	OriginalValue *string     `json:"original_value,omitempty"`
	ISO639_1      *string     `json:"iso_639_1,omitempty"`
}

type ValueClass struct {
	EpisodeID     int `json:"episode_id"`
	EpisodeNumber int `json:"episode_number"`
}

type ValueUnion struct {
	String     *string
	ValueClass *ValueClass
}

func (x *ValueUnion) UnmarshalJSON(data []byte) error {
	x.ValueClass = nil
	var c ValueClass
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.ValueClass = &c
	}
	return nil
}

func (x *ValueUnion) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.ValueClass != nil, x.ValueClass, false, nil, false, nil, false)
}

func unmarshalUnion(data []byte, pi **int, pf **float64, pb **bool, ps **string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) (bool, error) {
	if pi != nil {
		*pi = nil
	}
	if pf != nil {
		*pf = nil
	}
	if pb != nil {
		*pb = nil
	}
	if ps != nil {
		*ps = nil
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	tok, err := dec.Token()
	if err != nil {
		return false, err
	}

	switch v := tok.(type) {
	case json.Number:
		if pi != nil {
			i, err := v.Int64()
			if err == nil {
				a := int(i)
				*pi = &a
				return false, nil
			}
		}
		if pf != nil {
			f, err := v.Float64()
			if err == nil {
				*pf = &f
				return false, nil
			}
			return false, errors.New("Unparsable number")
		}
		return false, errors.New("Union does not contain number")
	case float64:
		return false, errors.New("Decoder should not return float64")
	case bool:
		if pb != nil {
			*pb = &v
			return false, nil
		}
		return false, errors.New("Union does not contain bool")
	case string:
		if haveEnum {
			return false, json.Unmarshal(data, pe)
		}
		if ps != nil {
			*ps = &v
			return false, nil
		}
		return false, errors.New("Union does not contain string")
	case nil:
		if nullable {
			return false, nil
		}
		return false, errors.New("Union does not contain null")
	case json.Delim:
		if v == '{' {
			if haveObject {
				return true, json.Unmarshal(data, pc)
			}
			if haveMap {
				return false, json.Unmarshal(data, pm)
			}
			return false, errors.New("Union does not contain object")
		}
		if v == '[' {
			if haveArray {
				return false, json.Unmarshal(data, pa)
			}
			return false, errors.New("Union does not contain array")
		}
		return false, errors.New("Cannot handle delimiter")
	}
	return false, errors.New("Cannot unmarshal union")

}
