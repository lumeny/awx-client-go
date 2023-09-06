// This file contains the implementation of the resource that manages the collection of
// workflow job templates.

package awx

import (
	"fmt"

	"github.com/CenturyLink/hca-awx-client-go/awx/internal/data"
)

type WorkflowJobTemplatesResource struct {
	Resource
}

func NewWorkflowJobTemplatesResource(connection *Connection, path string) *WorkflowJobTemplatesResource {
	resource := new(WorkflowJobTemplatesResource)
	resource.connection = connection
	resource.path = path
	return resource
}

func (r *WorkflowJobTemplatesResource) Get() *WorkflowJobTemplatesGetRequest {
	request := new(WorkflowJobTemplatesGetRequest)
	request.resource = &r.Resource
	return request
}

func (r *WorkflowJobTemplatesResource) Id(id int) *WorkflowJobTemplateResource {
	return NewWorkflowJobTemplateResource(r.connection, fmt.Sprintf("%s/%d", r.path, id))
}

type WorkflowJobTemplatesGetRequest struct {
	Request
}

func (r *WorkflowJobTemplatesGetRequest) Filter(name string, value interface{}) *WorkflowJobTemplatesGetRequest {
	r.addFilter(name, value)
	return r
}

func (r *WorkflowJobTemplatesGetRequest) Send() (response *WorkflowJobTemplatesGetResponse, err error) {
	output := new(data.WorkflowJobTemplatesGetResponse)
	err = r.get(output)
	if err != nil {
		return
	}
	response = new(WorkflowJobTemplatesGetResponse)
	response.count = output.Count
	response.previous = output.Previous
	response.next = output.Next
	response.results = make([]*WorkflowJobTemplate, len(output.Results))
	for i := 0; i < len(output.Results); i++ {
		response.results[i] = new(WorkflowJobTemplate)
		response.results[i].id = output.Results[i].Id
		response.results[i].name = output.Results[i].Name
		response.results[i].askLimitOnLaunch = output.Results[i].AskLimitOnLaunch
		response.results[i].askVarsOnLaunch = output.Results[i].AskVarsOnLaunch
	}
	return
}

type WorkflowJobTemplatesGetResponse struct {
	ListGetResponse

	results []*WorkflowJobTemplate
}

func (r *WorkflowJobTemplatesGetResponse) Results() []*WorkflowJobTemplate {
	return r.results
}
