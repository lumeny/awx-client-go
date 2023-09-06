// This file contains the data structures used to receive lists of job templates.

package data

type JobsGetResponse struct {
	ListGetResponse

	Results []*Job `json:"results,omitempty"`
}
