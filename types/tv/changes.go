package tv

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
	ID            string              `json:"id"`
	Action        string              `json:"action"`
	Time          string              `json:"time"`
	Value         *ValueUnion         `json:"value"`
	ISO639_1      *string             `json:"iso_639_1,omitempty"`
	OriginalValue *OriginalValueUnion `json:"original_value"`
}

type OriginalValueClass struct {
	ID         *int    `json:"id,omitempty"`
	Name       *string `json:"name,omitempty"`
	CreditID   *string `json:"credit_id,omitempty"`
	PersonID   *int    `json:"person_id,omitempty"`
	SeasonID   *int    `json:"season_id,omitempty"`
	Poster     *Poster `json:"poster,omitempty"`
	Department *string `json:"department,omitempty"`
	Job        *string `json:"job,omitempty"`
}

type Poster struct {
	FilePath string  `json:"file_path"`
	ISO639_1 *string `json:"iso_639_1"`
}

type ValueClass struct {
	SeasonID         *int    `json:"season_id,omitempty"`
	SeasonNumber     *int    `json:"season_number,omitempty"`
	ID               *int    `json:"id,omitempty"`
	Name             *string `json:"name,omitempty"`
	AddToEverySeason *bool   `json:"add_to_every_season,omitempty"`
	Character        *string `json:"character,omitempty"`
	CreditID         *string `json:"credit_id,omitempty"`
	Order            *int    `json:"order,omitempty"`
	PersonID         *int    `json:"person_id,omitempty"`
	Poster           *Poster `json:"poster,omitempty"`
	Department       *string `json:"department,omitempty"`
	Job              *string `json:"job,omitempty"`
}

type OriginalValueUnion struct {
	OriginalValueClass *OriginalValueClass
	String             *string
}

func (x *OriginalValueUnion) UnmarshalJSON(data []byte) error {
	x.OriginalValueClass = nil
	var c OriginalValueClass
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.OriginalValueClass = &c
	}
	return nil
}

func (x *OriginalValueUnion) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.OriginalValueClass != nil, x.OriginalValueClass, false, nil, false, nil, false)
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

func marshalUnion(pi *int, pf *float64, pb *bool, ps *string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) ([]byte, error) {
	if pi != nil {
		return json.Marshal(*pi)
	}
	if pf != nil {
		return json.Marshal(*pf)
	}
	if pb != nil {
		return json.Marshal(*pb)
	}
	if ps != nil {
		return json.Marshal(*ps)
	}
	if haveArray {
		return json.Marshal(pa)
	}
	if haveObject {
		return json.Marshal(pc)
	}
	if haveMap {
		return json.Marshal(pm)
	}
	if haveEnum {
		return json.Marshal(pe)
	}
	if nullable {
		return json.Marshal(nil)
	}
	return nil, errors.New("Union must not be null")
}
