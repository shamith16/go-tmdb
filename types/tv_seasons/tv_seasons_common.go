package tv_seasons

import (
	"encoding/json"
	"errors"
)

type Result struct {
	ISO639_1      string      `json:"iso_639_1,omitempty"`
	ISO3166_1     string      `json:"iso_3166_1,omitempty"`
	Name          string      `json:"name,omitempty"`
	Key           string      `json:"key,omitempty"`
	Site          string      `json:"site,omitempty"`
	Size          int         `json:"size,omitempty"`
	Type          string      `json:"type,omitempty"`
	Official      bool        `json:"official,omitempty"`
	PublishedAt   string      `json:"published_at,omitempty"`
	ID            string      `json:"id,omitempty"`
	EpisodeNumber int         `json:"episode_number,omitempty"`
	Rated         *RatedUnion `json:"rated,omitempty"`
}

type Cast struct {
	Adult              bool    `json:"adult,omitempty"`
	Gender             int     `json:"gender,omitempty"`
	ID                 int     `json:"id,omitempty"`
	KnownForDepartment string  `json:"known_for_department,omitempty"`
	Name               string  `json:"name,omitempty"`
	OriginalName       string  `json:"original_name,omitempty"`
	Popularity         float64 `json:"popularity,omitempty"`
	ProfilePath        *string `json:"profile_path,omitempty"`
	Character          *string `json:"character,omitempty,omitempty"`
	CreditID           string  `json:"credit_id,omitempty"`
	Order              *int    `json:"order,omitempty,omitempty"`
	Department         *string `json:"department,omitempty,omitempty"`
	Job                *string `json:"job,omitempty,omitempty"`
	Roles              []Role  `json:"roles,omitempty,omitempty"`
	TotalEpisodeCount  int     `json:"total_episode_count,omitempty"`
	Jobs               []Job   `json:"jobs,omitempty,omitempty"`
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
