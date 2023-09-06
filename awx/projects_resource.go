// This file contains the implementation of the resource that manages the collection of
// projects.

package awx

import (
	"fmt"

	"github.com/CenturyLink/hca-awx-client-go/awx/internal/data"
)

type ProjectsResource struct {
	Resource
}

func NewProjectsResource(connection *Connection, path string) *ProjectsResource {
	resource := new(ProjectsResource)
	resource.connection = connection
	resource.path = path
	return resource
}

func (r *ProjectsResource) Get() *ProjectsGetRequest {
	request := new(ProjectsGetRequest)
	request.resource = &r.Resource
	return request
}

func (r *ProjectsResource) Id(id int) *ProjectResource {
	return NewProjectResource(r.connection, fmt.Sprintf("%s/%d", r.path, id))
}

type ProjectsGetRequest struct {
	Request
}

func (r *ProjectsGetRequest) Filter(name string, value interface{}) *ProjectsGetRequest {
	r.addFilter(name, value)
	return r
}

func (r *ProjectsGetRequest) Send() (response *ProjectsGetResponse, err error) {
	output := new(data.ProjectsGetResponse)
	err = r.get(output)
	if err != nil {
		return
	}
	response = new(ProjectsGetResponse)
	response.count = output.Count
	response.previous = output.Previous
	response.next = output.Next
	response.results = make([]*Project, len(output.Results))
	for i := 0; i < len(output.Results); i++ {
		response.results[i] = new(Project)
		response.results[i].id = output.Results[i].Id
		response.results[i].name = output.Results[i].Name
		response.results[i].scmType = output.Results[i].SCMType
		response.results[i].scmURL = output.Results[i].SCMURL
		response.results[i].scmBranch = output.Results[i].SCMBranch
	}
	return
}

type ProjectsGetResponse struct {
	ListGetResponse

	results []*Project
}

func (r *ProjectsGetResponse) Results() []*Project {
	return r.results
}
