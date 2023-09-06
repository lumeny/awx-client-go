// This file contains the implementation of the job template type.

package awx

type JobStatus string

const (
	JobStatusNew       JobStatus = "new"
	JobStatusPending   JobStatus = "pending"
	JobStatusWaiting   JobStatus = "waiting"
	JobStatusRunning   JobStatus = "running"
	JobStatusSuccesful JobStatus = "successful"
	JobStatusFailed    JobStatus = "failed"
	JobStatusError     JobStatus = "error"
	JobStatusCancelled JobStatus = "canceled"
)

type Job struct {
	id     int
	status JobStatus
}

func (j *Job) Id() int {
	return j.id
}

func (j *Job) Status() JobStatus {
	return j.status
}

func (j *Job) IsFinished() bool {
	switch j.status {
	case
		JobStatusSuccesful,
		JobStatusFailed,
		JobStatusError,
		JobStatusCancelled:
		return true
	}
	return false
}

func (j *Job) IsSuccessful() bool {
	return j.status == JobStatusSuccesful
}
