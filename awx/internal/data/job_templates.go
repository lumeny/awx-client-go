// This file contains the data structures used to receive lists of job templates.

package data

type JobTemplatesGetResponse struct {
	ListGetResponse

	Results []*JobTemplate `json:"results,omitempty"`
}
