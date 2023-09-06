// This file contains the implementation of the resource that manages a specific workflow job
// template.

package awx

import (
	"github.com/CenturyLink/hca-awx-client-go/awx/internal/data"
)

type WorkflowJobTemplateResource struct {
	Resource
}

func NewWorkflowJobTemplateResource(connection *Connection, path string) *WorkflowJobTemplateResource {
	resource := new(WorkflowJobTemplateResource)
	resource.connection = connection
	resource.path = path
	return resource
}

func (r *WorkflowJobTemplateResource) Get() *WorkflowJobTemplateGetRequest {
	request := new(WorkflowJobTemplateGetRequest)
	request.resource = &r.Resource
	return request
}

func (r *WorkflowJobTemplateResource) Launch() *WorkflowJobTemplateLaunchResource {
	return NewWorkflowJobTemplateLaunchResource(r.connection, r.path+"/launch")
}

type WorkflowJobTemplateGetRequest struct {
	Request
}

func (r *WorkflowJobTemplateGetRequest) Send() (response *WorkflowJobTemplateGetResponse, err error) {
	output := new(data.WorkflowJobTemplateGetResponse)
	err = r.get(output)
	if err != nil {
		return
	}
	response = new(WorkflowJobTemplateGetResponse)
	response.result = new(WorkflowJobTemplate)
	response.result.id = output.Id
	response.result.name = output.Name
	response.result.askLimitOnLaunch = output.AskLimitOnLaunch
	response.result.askVarsOnLaunch = output.AskVarsOnLaunch

	return
}

type WorkflowJobTemplateGetResponse struct {
	result *WorkflowJobTemplate
}

func (r *WorkflowJobTemplateGetResponse) Result() *WorkflowJobTemplate {
	return r.result
}
