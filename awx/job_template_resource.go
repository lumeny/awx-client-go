// This file contains the implementation of the resource that manages a specific job
// template.

package awx

import (
	"github.com/CenturyLink/hca-awx-client-go/awx/internal/data"
)

type JobTemplateResource struct {
	Resource
}

func NewJobTemplateResource(connection *Connection, path string) *JobTemplateResource {
	resource := new(JobTemplateResource)
	resource.connection = connection
	resource.path = path
	return resource
}

func (r *JobTemplateResource) Get() *JobTemplateGetRequest {
	request := new(JobTemplateGetRequest)
	request.resource = &r.Resource
	return request
}

func (r *JobTemplateResource) Launch() *JobTemplateLaunchResource {
	return NewJobTemplateLaunchResource(r.connection, r.path+"/launch")
}

type JobTemplateGetRequest struct {
	Request
}

func (r *JobTemplateGetRequest) Send() (response *JobTemplateGetResponse, err error) {
	output := new(data.JobTemplateGetResponse)
	err = r.get(output)
	if err != nil {
		return
	}
	response = new(JobTemplateGetResponse)
	response.result = new(JobTemplate)
	response.result.id = output.Id
	response.result.name = output.Name
	response.result.askLimitOnLaunch = output.AskLimitOnLaunch
	response.result.askVarsOnLaunch = output.AskVarsOnLaunch

	return
}

type JobTemplateGetResponse struct {
	result *JobTemplate
}

func (r *JobTemplateGetResponse) Result() *JobTemplate {
	return r.result
}
