// This file contains the data structures used to receive lists of job templates.

package data

type ProjectsGetResponse struct {
	ListGetResponse

	Results []*Project `json:"results,omitempty"`
}
