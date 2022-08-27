package configuration

import "encoding/json"

type Jobs []Job

func UnmarshalJobs(data []byte) (Jobs, error) {
	var r Jobs
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Jobs) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Job struct {
	Department string   `json:"department"`
	Jobs       []string `json:"jobs"`
}
