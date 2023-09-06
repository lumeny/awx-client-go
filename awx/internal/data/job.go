// This file contains the data structures used for sending and receiving jobs.

package data

type Job struct {
	Id     int    `json:"id,omitempty"`
	Status string `json:"status,omitempty"`
}

type JobGetResponse struct {
	Job
}
