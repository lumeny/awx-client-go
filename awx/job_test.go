// This file contains the implementation of the job template type.

package awx

import "testing"

func TestIsSuccessful(t *testing.T) {
	for _, status := range []JobStatus{
		JobStatusNew, JobStatusPending, JobStatusWaiting, JobStatusRunning,
		JobStatusFailed, JobStatusError, JobStatusCancelled,
	} {
		if (&Job{0, status}).IsSuccessful() {
			t.Errorf("Job.IsSuccessful() Should return false for %s", status)
		}
	}
	if !(&Job{0, JobStatusSuccesful}).IsSuccessful() {
		t.Errorf("Job.IsSuccessful() Should return true for JobStatusSuccesful")
	}
}

func TestIsFinished(t *testing.T) {
	for _, status := range []JobStatus{
		JobStatusNew, JobStatusPending, JobStatusWaiting, JobStatusRunning,
	} {
		if (&Job{0, status}).IsFinished() {
			t.Errorf("Job.IsFinished() Should return false for %s", status)
		}
	}
	for _, status := range []JobStatus{
		JobStatusSuccesful, JobStatusFailed, JobStatusError, JobStatusCancelled,
	} {
		if !(&Job{0, status}).IsFinished() {
			t.Errorf("Job.IsFinished() Should return false for %s", status)
		}
	}
}
