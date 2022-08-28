package tv

import "encoding/json"

func UnmarshalAggregateCredits(data []byte) (AggregateCredits, error) {
	var r AggregateCredits
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AggregateCredits) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AggregateCredits struct {
	Cast []Cast `json:"cast"`
	Crew []Cast `json:"crew"`
	ID   int    `json:"id"`
}

type Job struct {
	CreditID     string `json:"credit_id"`
	Job          string `json:"job"`
	EpisodeCount int    `json:"episode_count"`
}

type Role struct {
	CreditID     string `json:"credit_id"`
	Character    string `json:"character"`
	EpisodeCount int    `json:"episode_count"`
}
