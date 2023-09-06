// This file contains the implementation of the resource that manages a specific project.

package awx

import (
	"github.com/CenturyLink/hca-awx-client-go/awx/internal/data"
)

type ProjectResource struct {
	Resource
}

func NewProjectResource(connection *Connection, path string) *ProjectResource {
	resource := new(ProjectResource)
	resource.connection = connection
	resource.path = path
	return resource
}

func (r *ProjectResource) Get() *ProjectGetRequest {
	request := new(ProjectGetRequest)
	request.resource = &r.Resource
	return request
}

type ProjectGetRequest struct {
	Request
}

func (r *ProjectGetRequest) Send() (response *ProjectGetResponse, err error) {
	output := new(data.ProjectGetResponse)
	err = r.get(output)
	if err != nil {
		return
	}
	response = new(ProjectGetResponse)
	response.result = new(Project)
	response.result.id = output.Id
	response.result.name = output.Name
	response.result.scmType = output.SCMType
	response.result.scmURL = output.SCMURL
	response.result.scmBranch = output.SCMBranch
	return
}

type ProjectGetResponse struct {
	result *Project
}

func (r *ProjectGetResponse) Result() *Project {
	return r.result
}
