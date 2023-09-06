// This file contains the implementation of the resource that manages launching of jobs from job
// templates.

package awx

import (
	"github.com/CenturyLink/hca-awx-client-go/awx/internal/data"
)

type JobResource struct {
	Resource
}

func NewJobResource(connection *Connection, path string) *JobResource {
	resource := new(JobResource)
	resource.connection = connection
	resource.path = path
	return resource
}

func (r *JobResource) Get() *JobGetRequest {
	request := new(JobGetRequest)
	request.resource = &r.Resource
	return request
}

type JobGetRequest struct {
	Request
}

func (r *JobGetRequest) Send() (response *JobGetResponse, err error) {
	output := new(data.JobGetResponse)
	err = r.get(output)
	if err != nil {
		return nil, err
	}
	response = new(JobGetResponse)
	if output != nil {
		response.job = new(Job)
		response.job.id = output.Id
		response.job.status = (JobStatus)(output.Status)
	}
	return
}

type JobGetResponse struct {
	job *Job
}

func (r *JobGetResponse) Job() *Job {
	return r.job
}
